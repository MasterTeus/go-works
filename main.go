package main

import (
	"github.com/MasterTeus/go-works.git/config"
	"github.com/MasterTeus/go-works.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization erro: %v", err)
		return
	}

	// initialize router
	router.Initalize()
}
