package main

import (
	"os"

	"github.com/kot-zakhar/golang_pet/internal/config"
	"github.com/kot-zakhar/golang_pet/internal/core"
)

func main() {
	port := ":80"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	config := config.GetConfig()

	core.InitializeDI(config)

	core.SetUpServer(port)
}
