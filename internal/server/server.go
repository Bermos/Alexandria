package server

import (
	"Alexandria/internal/gen/openapi"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Start(port int) {
	modulesApiService := NewModulesApiService()
	modulesApiController := openapi.NewModulesApiController(modulesApiService)

	terraformApiService := NewTerraformApiService()
	terraformApiController := openapi.NewTerraformApiController(terraformApiService)

	router := openapi.NewRouter(modulesApiController, terraformApiController)

	// Redirect requests to '/.well-known/terraform.json' to '/api/.well-known/terraform.json'
	router.Path("/.well-known/terraform.json").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api/.well-known/terraform.json", http.StatusMovedPermanently)
	})

	// Serve static frontend page
	fs := http.Dir("web/dist/web")
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := fs.Open(r.URL.Path[1:])
		if err != nil || r.URL.Path[1:] == "" {
			http.ServeFile(w, r, "web/dist/web/index.html")
		} else {
			http.ServeContent(w, r, r.URL.Path[1:], time.Now(), file)
		}
	})

	log.Printf("INFO - Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
