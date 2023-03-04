package terraform_bind

import (
	ctx "context"
	"github.com/hashicorp/terraform-exec/tfexec"
	"log"
	"os"
	"path/filepath"
)

type Module struct {
	Path       string
	Name       string
	HasReadme  bool
	ModuleInfo ModuleInfo
}

type ModuleInfo struct {
	Valid        bool
	ErrorCount   int
	WarningCount int
}

var (
	loadedModules = make(map[string]Module)
)

func Load(path string) {
	moduleDir, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("ERROR - error reading module path: %v", err)
	}

	moduleCh := make(chan Module, len(moduleDir))

	for _, entry := range moduleDir {
		if !entry.IsDir() {
			continue
		}

		go func(entry os.DirEntry) {

			modulePath := filepath.Join(path, entry.Name())
			log.Printf("INFO - Loading module at path %s", modulePath)

			_, fileReadErr := os.ReadFile(filepath.Join(modulePath, "README.md"))

			moduleCh <- Module{
				Path:       modulePath,
				Name:       entry.Name(),
				HasReadme:  fileReadErr == nil,
				ModuleInfo: getModuleInfo(modulePath),
			}
		}(entry)
	}

	for range moduleDir {
		module := <-moduleCh
		loadedModules[module.Name] = module
	}

	log.Printf("Loaded modules - %v", loadedModules)
}

func GetModules() map[string]Module {
	return loadedModules
}

func getModuleInfo(modulePath string) (moduleInfo ModuleInfo) {
	log.Printf("INFO - Validating module at path %s", modulePath)

	// Create tmp directory and defer its removal for running terraform init in a clean env
	tmpDir, err := os.MkdirTemp("", "tf_module_validate-")
	if err != nil {
		log.Fatalf("error running os.MkdirTemp: %s", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create tfexec.Terraform with a tf executable and in a workdir
	tf, err := tfexec.NewTerraform(tmpDir, tfExecPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	// Run terraform init in the tmp dir with the cloned source module
	err = tf.Init(ctx.Background(), tfexec.FromModule(modulePath))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	// Run terraform validate and store its results for moduleInfo
	validate, err := tf.Validate(ctx.Background())
	if err != nil {
		log.Fatalf("error running Validate: %s", err)
	}
	moduleInfo.Valid = validate.Valid
	moduleInfo.ErrorCount = validate.ErrorCount
	moduleInfo.WarningCount = validate.WarningCount

	// Fetch the required providers. TODO: parse the data and add to moduleInfo
	_, err = tf.ProvidersSchema(ctx.Background())
	if err != nil {
		log.Fatalf("error running Providers: %s", err)
	}

	// Fetch the state. TODO: parse the data and add to moduleInfo
	_, err = tf.Show(ctx.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	return
}
