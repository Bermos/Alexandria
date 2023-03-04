# Alexandria

This is a tool that allows users to store and deploy Terraform modules and providers. With this tool, users can easily manage their Terraform infrastructure code and deploy individual modules from a pre-defined catalog.

Features:
- Store Terraform modules and providers in a central repository
- Easily search and find modules by name or description
- Deploy individual modules from a pre-defined catalog
- Automatically download and install provider dependencies for each module
- View information about each module, including its source, version, and dependencies
- Manage access to modules with role-based access control

## Roadmap
implements:
- https://developer.hashicorp.com/terraform/internals/login-protocol
- https://developer.hashicorp.com/terraform/internals/module-registry-protocol
- https://developer.hashicorp.com/terraform/internals/provider-registry-protocol

## Development

### Setup
- (Re-)Generate server code from OpenAPI spec; `docker run --rm -v ($pwd):/local --name openapi-generator -u 1000 -w /local openapitools/openapi-generator-cli:latest generate -i openapi/spec.yaml -o internal/gen -g go-server -c openapi/server-config.yml`
- (Re-)Generate client code from OpenAPI spec: `docker run --rm -v ($pwd):/local --name openapi-generator -u 1000 -w /local openapitools/openapi-generator-cli:latest generate -i openapi/spec.yaml -o web/src/app/api_client -g typescript-angular`
