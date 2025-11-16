package router

import (
	"exptracker/internal/app"
	"exptracker/internal/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg config.Config) {
	app := app.NewApplication(cfg)
	v1 := r.Group("/api/v1")

	RegisterAuthRoutes(v1, app.AuthController)
}
