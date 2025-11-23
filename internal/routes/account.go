package router

import (
	accctrl "exptracker/internal/controller/account"

	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(rg *gin.RouterGroup, c *accctrl.Controller) {
	acc := rg.Group("/accounts")

	acc.POST("/", c.Create)
	acc.GET("/", c.GetAll)
	acc.GET("/:id", c.GetByID)
	acc.PUT("/:id", c.Update)
	acc.DELETE("/:id", c.Delete)
}
