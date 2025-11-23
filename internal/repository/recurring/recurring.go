package recurring

import (
	"exptracker/internal/database"
	models "exptracker/internal/domain/model"
	"time"
)

type Repository interface {
	Create(r *models.RecurringExpense) error
	GetAll(userID string) ([]models.RecurringExpense, error)
	GetByID(id string, userID string) (*models.RecurringExpense, error)
	Update(r *models.RecurringExpense) error
	Delete(id, userID string) error
	GetDue(now time.Time) ([]models.RecurringExpense, error)
}

type recurringRepository struct{}

func NewRecurringRepository() Repository {
	return &recurringRepository{}
}

func (r *recurringRepository) Create(e *models.RecurringExpense) error {
	return database.DB.Create(e).Error
}

func (r *recurringRepository) GetAll(userID string) ([]models.RecurringExpense, error) {
	var list []models.RecurringExpense
	err := database.DB.Where("user_id = ?", userID).Order("next_run_at").Find(&list).Error
	return list, err
}

func (r *recurringRepository) GetByID(id, userID string) (*models.RecurringExpense, error) {
	var e models.RecurringExpense
	err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&e).Error
	return &e, err
}

func (r *recurringRepository) Update(e *models.RecurringExpense) error {
	return database.DB.Save(e).Error
}

func (r *recurringRepository) Delete(id, userID string) error {
	return database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.RecurringExpense{}).Error
}

func (r *recurringRepository) GetDue(now time.Time) ([]models.RecurringExpense, error) {
	var list []models.RecurringExpense
	err := database.DB.Where("next_run_at <= ?", now).Find(&list).Error
	return list, err
}
