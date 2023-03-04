package server

import (
	"context"
	"github.com/Bermos/Alexandria/internal/gen/openapi"
	"github.com/Bermos/Alexandria/internal/terraform_services"
	"net/http"
)

type TerraformApiService struct {
}

// NewTerraformApiService creates a default api service
func NewTerraformApiService() openapi.TerraformApiServicer {
	return &TerraformApiService{}
}

func (s *TerraformApiService) WellKnownTerraformJsonGet(_ context.Context) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusOK, terraform_services.GetEnabledServices()), nil
}

func (s *TerraformApiService) TerraformModulesV1NamespaceNameSystemVersionDownloadGet(ctx context.Context, namespace string, name string, system string, version string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TerraformApiService) TerraformModulesV1NamespaceNameSystemVersionsGet(ctx context.Context, namespace string, name string, system string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TerraformApiService) TerraformProvidersV1NamespaceTypeVersionDownloadOsArchGet(ctx context.Context, namespace string, pType string, version string, os string, arch string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *TerraformApiService) TerraformProvidersV1NamespaceTypeVersionsGet(ctx context.Context, namespace string, pType string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
