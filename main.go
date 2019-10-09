package main

import (
	"fmt"
	"github.com/gin-rest-prisma-boilerplate/config"
)

func main() {
	// set this when moving between production and development configs
	yaml := config.Configuration("development")
	fmt.Print(yaml.Get("server.port"))
}
