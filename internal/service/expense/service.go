package expense

import (
	dto "exptracker/internal/domain/dto/expense"
	models "exptracker/internal/domain/model"
	repository "exptracker/internal/repository/expense"
	audit "exptracker/internal/service/audit"
)

type Service interface {
	Create(req dto.CreateExpenseRequest, userID string) (*models.Expense, error)
	GetAll(userID string) ([]models.Expense, error)
	GetByID(id, userID string) (*models.Expense, error)
	Update(id string, req dto.UpdateExpenseRequest, userID string) (*models.Expense, error)
	Delete(id, userID string) error
}

type expenseService struct {
	repo  repository.Repository
	audit audit.Service
}

func NewExpenseService(repo repository.Repository, auditsvc audit.Service) Service {
	return &expenseService{repo, auditsvc}
}

func (s *expenseService) Create(req dto.CreateExpenseRequest, userID string) (*models.Expense, error) {
	exp := &models.Expense{
		UserID:   userID,
		Amount:   req.Amount,
		Category: req.Category,
		Note:     req.Note,
	}
	s.audit.Log(userID, "create", "expense", exp.ID, exp)
	return exp, s.repo.Create(exp)
}

func (s *expenseService) GetAll(userID string) ([]models.Expense, error) {
	return s.repo.GetAll(userID)
}

func (s *expenseService) GetByID(id, userID string) (*models.Expense, error) {
	return s.repo.GetByID(id, userID)
}

func (s *expenseService) Update(id string, req dto.UpdateExpenseRequest, userID string) (*models.Expense, error) {
	exp, err := s.repo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	if req.Amount != 0 {
		exp.Amount = req.Amount
	}
	if req.Category != "" {
		exp.Category = req.Category
	}
	exp.Note = req.Note
	s.audit.Log(userID, "update", "expense", exp.ID, exp)
	return exp, s.repo.Update(exp)
}

func (s *expenseService) Delete(id, userID string) error {
	s.audit.Log(userID, "delete", "expense", id, nil)
	return s.repo.Delete(id, userID)
}
