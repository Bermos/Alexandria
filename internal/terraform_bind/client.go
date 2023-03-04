package terraform_bind

import (
	"context"
	"github.com/hashicorp/go-version"
	"log"

	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
)

var tfExecPath string

func Init() {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.3.0")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}

	tfExecPath = execPath
}
