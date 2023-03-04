package main

import (
	"github.com/Bermos/Alexandria/internal/server"
	"github.com/Bermos/Alexandria/internal/terraform_bind"
)

func main() {
	terraform_bind.Init()
	terraform_bind.Load("")

	server.Start(9090)
}
