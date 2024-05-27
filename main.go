package main

import (
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/config"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("error ocurred on main: %v", err)
		return
	}

	router.Initialize()
}
