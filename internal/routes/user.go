package router

import (
	userctrl "exptracker/internal/controller/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, c *userctrl.Controller) {
	usr := rg.Group("/user")

	usr.GET("/me", c.GetProfile)
	usr.PUT("/update", c.Update)
	usr.PUT("/change-password", c.ChangePassword)
	usr.DELETE("/delete", c.Delete)
}
