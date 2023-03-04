package server

import (
	"context"
	"github.com/Bermos/Alexandria/internal/gen/openapi"
	"github.com/Bermos/Alexandria/internal/terraform_bind"
	"net/http"
)

type ModulesApiService struct {
}

// NewModulesApiService creates a default api service
func NewModulesApiService() openapi.ModulesApiServicer {
	return &ModulesApiService{}
}

func (s *ModulesApiService) ModulesGet(_ context.Context) (openapi.ImplResponse, error) {
	resp := make(map[string]openapi.Module)

	for k, v := range terraform_bind.GetModules() {

		resp[k] = openapi.Module{
			Path:      v.Path,
			Name:      v.Name,
			HasReadme: v.HasReadme,
			ModuleInfo: openapi.ModuleModuleInfo{
				Valid:        v.ModuleInfo.Valid,
				ErrorCount:   int32(v.ModuleInfo.ErrorCount),
				WarningCount: int32(v.ModuleInfo.WarningCount),
			},
		}
	}

	return openapi.Response(http.StatusOK, resp), nil
}
