package user

import (
	"errors"

	dto "exptracker/internal/domain/dto/user"
	models "exptracker/internal/domain/model"
	"exptracker/internal/repository/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetProfile(userID string) (*models.User, error)
	Update(userID string, req dto.UpdateUserRequest) (*models.User, error)
	ChangePassword(userID string, req dto.ChangePasswordRequest) error
	Delete(userID string) error
}

type userService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) Service {
	return &userService{repo}
}

func (s *userService) GetProfile(userID string) (*models.User, error) {
	return s.repo.GetByID(userID)
}

func (s *userService) Update(userID string, req dto.UpdateUserRequest) (*models.User, error) {
	u, err := s.repo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if req.Role != "" {
		u.Role = req.Role
	}

	return u, s.repo.Update(u)
}

func (s *userService) ChangePassword(userID string, req dto.ChangePasswordRequest) error {
	u, err := s.repo.GetByID(userID)
	if err != nil {
		return err
	}

	// check old password
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.OldPassword)) != nil {
		return errors.New("incorrect old password")
	}

	// hash new password
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	u.Password = string(hash)

	return s.repo.Update(u)
}

func (s *userService) Delete(userID string) error {
	return s.repo.Delete(userID)
}
