package auth

import (
	"exptracker/internal/database"
	models "exptracker/internal/domain/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository() Repository {
	return &authRepository{
		db: database.DB,
	}
}

func (r *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *authRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
