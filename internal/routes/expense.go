package router

import (
	expctrl "exptracker/internal/controller/expense"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseRoutes(rg *gin.RouterGroup, c *expctrl.Controller) {
	exp := rg.Group("/expenses")

	exp.POST("/", c.Create)
	exp.GET("/", c.GetAll)
	exp.GET("/:id", c.GetByID)
	exp.PUT("/:id", c.Update)
	exp.DELETE("/:id", c.Delete)
}
