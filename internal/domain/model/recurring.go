package models

import "time"

type RecurringExpense struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	Amount    float64   `gorm:"type:numeric;not null" json:"amount"`
	Interval  string    `gorm:"type:varchar(20);not null"` // daily, weekly, monthly
	NextRunAt time.Time `json:"next_run_at"`
	Category  string    `gorm:"type:varchar(50)" json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
