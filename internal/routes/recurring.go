package router

import (
	recctrl "exptracker/internal/controller/recurring"

	"github.com/gin-gonic/gin"
)

func RegisterRecurringRoutes(rg *gin.RouterGroup, c *recctrl.Controller) {
	r := rg.Group("/recurring")

	r.POST("/", c.Create)
	r.GET("/", c.GetAll)
	r.PUT("/:id", c.Update)
	r.DELETE("/:id", c.Delete)
}
