package account

import (
	dto "exptracker/internal/domain/dto/account"
	models "exptracker/internal/domain/model"
	repo "exptracker/internal/repository/account"
)

type Service interface {
	Create(req dto.CreateAccountRequest, userID string) (*models.Account, error)
	GetAll(userID string) ([]models.Account, error)
	GetByID(id, userID string) (*models.Account, error)
	Update(id string, req dto.UpdateAccountRequest, userID string) (*models.Account, error)
	Delete(id, userID string) error
}

type accountService struct {
	repo repo.Repository
}

func NewAccountService(repo repo.Repository) Service {
	return &accountService{repo}
}

func (s *accountService) Create(req dto.CreateAccountRequest, userID string) (*models.Account, error) {
	acc := &models.Account{
		UserID:  userID,
		Name:    req.Name,
		Type:    req.Type,
		Balance: req.Balance,
	}

	return acc, s.repo.Create(acc)
}

func (s *accountService) GetAll(userID string) ([]models.Account, error) {
	return s.repo.GetAll(userID)
}

func (s *accountService) GetByID(id, userID string) (*models.Account, error) {
	return s.repo.GetByID(id, userID)
}

func (s *accountService) Update(id string, req dto.UpdateAccountRequest, userID string) (*models.Account, error) {
	acc, err := s.repo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		acc.Name = req.Name
	}
	if req.Type != "" {
		acc.Type = req.Type
	}
	acc.Balance = req.Balance

	return acc, s.repo.Update(acc)
}

func (s *accountService) Delete(id, userID string) error {
	return s.repo.Delete(id, userID)
}
