package main

import (
	"github.com/Archii0/lumo/pkg/api"
	"github.com/Archii0/lumo/pkg/logger"
)

func main() {
	logger.Init()
	router := api.SetupRouter()
	router.Run()
}
