package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RecurringExpense struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	Amount    float64   `gorm:"type:numeric;not null" json:"amount"`
	Interval  string    `gorm:"type:varchar(20);not null"`
	NextRunAt time.Time `json:"next_run_at"`
	Category  string    `gorm:"type:varchar(50)" json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

func (r *RecurringExpense) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.NewString()
	r.CreatedAt = time.Now()
	return nil
}
