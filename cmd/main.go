package main

import (
	"Alexandria/internal/server"
	"Alexandria/internal/terraform_bind"
)

func main() {
	terraform_bind.Init()
	terraform_bind.Load("")

	server.Start(9090)
}
