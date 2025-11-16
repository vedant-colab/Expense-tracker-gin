package app

import (
	"exptracker/internal/config"

	// Auth module imports
	authctrl "exptracker/internal/controller/auth"
	authrepo "exptracker/internal/repository/auth"
	authsvc "exptracker/internal/service/auth"
)

// Application is the central dependency container.
// All controllers live here (services + repositories injected).
type Application struct {
	AuthController *authctrl.Controller
	// Add more controllers here:
	// User *userctrl.Controller
	// Expense *expensectl.Controller
	// Account *accountctrl.Controller
}

// NewApplication builds all module dependencies (DI).
func NewApplication(cfg config.Config) *Application {

	// ---------------- AUTH MODULE ----------------
	authRepository := authrepo.NewAuthRepository()
	authService := authsvc.NewAuthService(authRepository, cfg)
	authController := authctrl.NewAuthController(authService)

	// --------------- ADD FUTURE MODULES HERE ----------------
	// userRepo := userrepo.NewUserRepository()
	// userService := usersvc.NewUserService(userRepo)
	// userController := userctrl.NewUserController(userService)

	// expenseRepo := expenserepo.NewExpenseRepository()
	// expenseService := expensesvc.NewExpenseService(expenseRepo)
	// expenseController := expensectrl.NewExpenseController(expenseService)

	return &Application{
		AuthController: authController,
		// User: userController,
		// Expense: expenseController,
		// Account: accountController,
	}
}
