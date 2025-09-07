package main

import (
	"github.com/Archii0/lumo/pkg/api"
)

func main() {
	router := api.SetupRouter()
	router.Run()
}
