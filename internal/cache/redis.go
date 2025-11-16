package cache

import (
	"context"
	"exptracker/internal/config"
	"exptracker/internal/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect(cfg config.Config) *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		logger.L().Fatal().Err(err).Msg("Failed to connect to Redis")
	}

	logger.L().Info().Msg("Connected to redis")
	return Client
}

func Close() {
	if Client == nil {
		return
	}
	if err := Client.Close(); err != nil {
		logger.L().Error().Err(err).Msg("Error closing Redis client")
	} else {
		logger.L().Info().Msg("Redis connection closed")
	}
}
