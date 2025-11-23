package expense

import (
	"exptracker/internal/database"
	models "exptracker/internal/domain/model"
)

type Repository interface {
	Create(expense *models.Expense) error
	GetByID(id string, userID string) (*models.Expense, error)
	GetAll(userID string) ([]models.Expense, error)
	Update(expense *models.Expense) error
	Delete(id string, userID string) error
}

type expenseRepository struct{}

func NewExpenseRepository() Repository {
	return &expenseRepository{}
}

func (r *expenseRepository) Create(expense *models.Expense) error {
	return database.DB.Create(expense).Error
}

func (r *expenseRepository) GetByID(id string, userID string) (*models.Expense, error) {
	var e models.Expense
	err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&e).Error
	return &e, err
}

func (r *expenseRepository) GetAll(userID string) ([]models.Expense, error) {
	var expenses []models.Expense
	err := database.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&expenses).Error
	return expenses, err
}

func (r *expenseRepository) Update(expense *models.Expense) error {
	return database.DB.Save(expense).Error
}

func (r *expenseRepository) Delete(id string, userID string) error {
	return database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Expense{}).Error
}
