package user

import (
	"exptracker/internal/database"
	models "exptracker/internal/domain/model"
)

type Repository interface {
	GetByID(id string) (*models.User, error)
	Update(u *models.User) error
	Delete(id string) error
	GetByEmail(email string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() Repository {
	return &userRepository{}
}

func (r *userRepository) GetByID(id string) (*models.User, error) {
	var u models.User
	err := database.DB.Where("id = ?", id).First(&u).Error
	return &u, err
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var u models.User
	err := database.DB.Where("email = ?", email).First(&u).Error
	return &u, err
}

func (r *userRepository) Update(u *models.User) error {
	return database.DB.Save(u).Error
}

func (r *userRepository) Delete(id string) error {
	return database.DB.Delete(&models.User{}, "id = ?", id).Error
}
