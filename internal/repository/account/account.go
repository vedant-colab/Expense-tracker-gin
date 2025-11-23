package account

import (
	"exptracker/internal/database"
	models "exptracker/internal/domain/model"
)

type Repository interface {
	Create(acc *models.Account) error
	GetByID(id, userID string) (*models.Account, error)
	GetAll(userID string) ([]models.Account, error)
	Update(acc *models.Account) error
	Delete(id, userID string) error
}

type accountRepository struct{}

func NewAccountRepository() Repository {
	return &accountRepository{}
}

func (r *accountRepository) Create(a *models.Account) error {
	return database.DB.Create(a).Error
}

func (r *accountRepository) GetByID(id, userID string) (*models.Account, error) {
	var a models.Account
	err := database.DB.Where("id = ? AND user_id = ?", id, userID).
		First(&a).Error
	return &a, err
}

func (r *accountRepository) GetAll(userID string) ([]models.Account, error) {
	var list []models.Account
	err := database.DB.Where("user_id = ?", userID).
		Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *accountRepository) Update(a *models.Account) error {
	return database.DB.Save(a).Error
}

func (r *accountRepository) Delete(id, userID string) error {
	return database.DB.Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.Account{}).Error
}
