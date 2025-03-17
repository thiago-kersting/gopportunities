package main

import (
	"github.com/thiago-kersting/gopportunities/config"
	"github.com/thiago-kersting/gopportunities/router"
)

var logger config.Logger

func main() {
	logger = *config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
