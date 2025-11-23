package main

import (
	app "exptracker/internal/app"
	"exptracker/internal/cache"
	"exptracker/internal/config"
	"exptracker/internal/database"
	"exptracker/internal/logger"
	scheduler "exptracker/internal/schedulers"
	"exptracker/internal/server"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg.App.Env)

	database.Connect(cfg)
	cache.Connect(cfg)

	scheduler := scheduler.New()
	app := app.NewApplication(cfg)
	scheduler.EveryMinute(app.RecurringService.Run)
	scheduler.Start()

	server.Run(cfg)
}
