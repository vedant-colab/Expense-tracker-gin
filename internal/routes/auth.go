package router

import (
	authctrl "exptracker/internal/controller/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, controller *authctrl.Controller) {
	auth := rg.Group("/auth")

	auth.POST("/register", controller.Register)
	auth.POST("/login", controller.Login)
	auth.POST("/refresh", controller.Refresh)
}
