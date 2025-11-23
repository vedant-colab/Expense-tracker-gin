package app

import (
	"exptracker/internal/config"

	authctrl "exptracker/internal/controller/auth"
	authrepo "exptracker/internal/repository/auth"
	authsvc "exptracker/internal/service/auth"

	expensectrl "exptracker/internal/controller/expense"
	expenserepo "exptracker/internal/repository/expense"
	expensesvc "exptracker/internal/service/expense"

	accountctrl "exptracker/internal/controller/account"
	accountrepo "exptracker/internal/repository/account"
	accountsvc "exptracker/internal/service/account"

	userctrl "exptracker/internal/controller/user"
	userrepo "exptracker/internal/repository/user"
	usersvc "exptracker/internal/service/user"

	recurringctrl "exptracker/internal/controller/recurring"
	recurringrepo "exptracker/internal/repository/recurring"
	recurringsvc "exptracker/internal/service/recurring"

	auditrepo "exptracker/internal/repository/audit"
	auditsvc "exptracker/internal/service/audit"
)

type Application struct {
	AuthController   *authctrl.Controller
	User             *userctrl.Controller
	Expense          *expensectrl.Controller
	Account          *accountctrl.Controller
	Recurring        *recurringctrl.Controller
	Audit            auditsvc.Service
	RecurringService recurringsvc.Service
}

func NewApplication(cfg config.Config) *Application {
	// ---------------- AUDIT MODULE ----------------
	auditRepo := auditrepo.NewAuditRepository()
	auditService := auditsvc.NewAuditService(auditRepo)

	// ---------------- AUTH MODULE ----------------
	authRepository := authrepo.NewAuthRepository()
	authService := authsvc.NewAuthService(authRepository, cfg)
	authController := authctrl.NewAuthController(authService)

	// ---------------- EXPENSE MODULE ----------------
	expenseRepo := expenserepo.NewExpenseRepository()
	expenseService := expensesvc.NewExpenseService(expenseRepo, auditService)
	expenseController := expensectrl.NewExpenseController(expenseService)

	// ---------------- ACCOUNT MODULE ----------------
	accountRepo := accountrepo.NewAccountRepository()
	accountService := accountsvc.NewAccountService(accountRepo)
	accountController := accountctrl.NewAccountController(accountService)

	// ---------------- USER MODULE ----------------
	userRepo := userrepo.NewUserRepository()
	userService := usersvc.NewUserService(userRepo)
	userController := userctrl.NewUserController(userService)

	// ---------------- RECURRING MODULE ----------------
	recRepo := recurringrepo.NewRecurringRepository()
	recService := recurringsvc.NewRecurringService(recRepo, expenseService)
	recController := recurringctrl.NewRecurringController(recService)

	return &Application{
		AuthController:   authController,
		User:             userController,
		Expense:          expenseController,
		Account:          accountController,
		Audit:            auditService,
		Recurring:        recController,
		RecurringService: recService,
	}
}
