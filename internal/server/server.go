package server

import (
	"context"
	"exptracker/internal/cache"
	"exptracker/internal/config"
	"exptracker/internal/database"
	"exptracker/internal/logger"
	appRouter "exptracker/internal/routes"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(cfg config.Config) {
	router := gin.Default()
	database.Connect(cfg)
	cache.Connect(cfg)
	appRouter.RegisterRoutes(router, cfg)

	addr := ":" + strconv.Itoa(cfg.Server.Port)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	logger.L().Info().
		Str("env", cfg.App.Env).
		Str("port", strconv.Itoa(cfg.Server.Port)).
		Msg("Starting expense tracker server")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.L().Fatal().Err(err).Msg("Server startup failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	logger.L().Warn().Msg("Shutdown signal recieved...Shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.L().Error().Err(err).Msg("Server forced to shutdown")
	}
	database.Close()
	cache.Close()
	logger.L().Info().Msg("Server exited cleanly")

}
