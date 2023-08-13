package main

import (
	"github.com/kot-zakhar/golang_pet/internal/config"
	"github.com/kot-zakhar/golang_pet/internal/core"
)

func main() {
	config := config.GetConfig()

	core.InitializeDI(config)

	core.SetUpServer(":80")
}
