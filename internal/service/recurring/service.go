package recurring

import (
	"errors"
	"time"

	expensedto "exptracker/internal/domain/dto/expense"
	dto "exptracker/internal/domain/dto/recurring"
	models "exptracker/internal/domain/model"
	repo "exptracker/internal/repository/recurring"
	expense "exptracker/internal/service/expense"
)

type Service interface {
	Create(req dto.CreateRecurringRequest, userID string) (*models.RecurringExpense, error)
	GetAll(userID string) ([]models.RecurringExpense, error)
	Update(id string, req dto.UpdateRecurringRequest, userID string) (*models.RecurringExpense, error)
	Delete(id, userID string) error

	RunDueExpenses() error
	Run()
}

type recurringService struct {
	repo        repo.Repository
	expenseServ expense.Service
}

func NewRecurringService(repo repo.Repository, exp expense.Service) Service {
	return &recurringService{repo, exp}
}

func isValidInterval(interval string) bool {
	switch interval {
	case "daily", "weekly", "monthly", "yearly":
		return true
	}
	return false
}

func nextRun(from time.Time, interval string) time.Time {
	switch interval {
	case "daily":
		return from.AddDate(0, 0, 1)
	case "weekly":
		return from.AddDate(0, 0, 7)
	case "monthly":
		return from.AddDate(0, 1, 0)
	case "yearly":
		return from.AddDate(1, 0, 0)
	}
	return from
}

func (s *recurringService) Create(req dto.CreateRecurringRequest, userID string) (*models.RecurringExpense, error) {

	if !isValidInterval(req.Interval) {
		return nil, errors.New("invalid interval")
	}

	r := &models.RecurringExpense{
		UserID:    userID,
		Amount:    req.Amount,
		Interval:  req.Interval,
		Category:  req.Category,
		NextRunAt: nextRun(time.Now(), req.Interval),
	}

	return r, s.repo.Create(r)
}

func (s *recurringService) GetAll(userID string) ([]models.RecurringExpense, error) {
	return s.repo.GetAll(userID)
}

func (s *recurringService) Update(id string, req dto.UpdateRecurringRequest, userID string) (*models.RecurringExpense, error) {
	r, err := s.repo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	if req.Amount != 0 {
		r.Amount = req.Amount
	}
	if req.Interval != "" {
		if !isValidInterval(req.Interval) {
			return nil, errors.New("invalid interval")
		}
		r.Interval = req.Interval
		r.NextRunAt = nextRun(time.Now(), req.Interval)
	}
	if req.Category != "" {
		r.Category = req.Category
	}

	return r, s.repo.Update(r)
}

func (s *recurringService) Delete(id, userID string) error {
	return s.repo.Delete(id, userID)
}

func (s *recurringService) RunDueExpenses() error {
	now := time.Now()

	dueItems, err := s.repo.GetDue(now)
	if err != nil {
		return err
	}

	for _, r := range dueItems {
		// Create actual Expense
		_, err := s.expenseServ.Create(
			expensedto.CreateExpenseRequest{
				Amount:   r.Amount,
				Category: r.Category,
				Note:     "Recurring",
			},
			r.UserID,
		)
		if err != nil {
			continue
		}

		// Update next run
		r.NextRunAt = nextRun(now, r.Interval)
		s.repo.Update(&r)
	}

	return nil
}

func (s *recurringService) Run() {
	s.RunDueExpenses()
}
