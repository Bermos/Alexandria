package server

import (
	"encoding/json"
	"github.com/Bermos/Alexandria/internal/terraform_services"
	"net/http"
)

type login struct {
	Client     string   `json:"client"`
	GrantTypes []string `json:"grant_types"`
	Authz      string   `json:"authz"`
	Token      string   `json:"token"`
	Ports      []int    `json:"ports"`
}

func handleWellKnown(w http.ResponseWriter, r *http.Request) {
	var wellKnownResponse map[string]interface{}
	for key, service := range terraform_services.GetEnabledServices() {
		wellKnownResponse[key] = service
	}

	json.NewEncoder(w).Encode(wellKnownResponse)
}
