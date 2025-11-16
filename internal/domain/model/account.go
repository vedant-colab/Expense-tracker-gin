package models

import "time"

type Account struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name"`
	Type      string    `gorm:"type:varchar(20);not null" json:"type"` // cash, bank, wallet, etc.
	Balance   float64   `gorm:"type:numeric" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
