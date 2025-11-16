package models

import "time"

type Expense struct {
	ID        string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string `gorm:"type:uuid;index;not null" json:"user_id"`
	AccountID string `gorm:"type:uuid;index;not null" json:"account_id"`

	Category string    `gorm:"type:varchar(50);index" json:"category"`
	Amount   float64   `gorm:"type:numeric;not null" json:"amount"`
	Note     string    `gorm:"type:text" json:"note"`
	Date     time.Time `json:"date"`

	CreatedAt time.Time `json:"created_at"`
}
