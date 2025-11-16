package main

import (
	"exptracker/internal/config"
	"exptracker/internal/logger"
	"exptracker/internal/server"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg.App.Env)

	server.Run(cfg)
}
