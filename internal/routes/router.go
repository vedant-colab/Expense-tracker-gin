package router

import (
	"exptracker/internal/app"
	"exptracker/internal/config"
	"exptracker/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg config.Config) {
	app := app.NewApplication(cfg)
	v1 := r.Group("/api/v1")

	protected := v1.Group("/")
	protected.Use(middleware.JWT(cfg))

	RegisterAuthRoutes(v1, app.AuthController)
	RegisterExpenseRoutes(protected, app.Expense)
	RegisterAccountRoutes(protected, app.Account)
	RegisterUserRoutes(protected, app.User)
	RegisterRecurringRoutes(protected, app.Recurring)
}
