package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id" gorm:"index"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.NewString()
	a.CreatedAt = time.Now()
	return nil
}
