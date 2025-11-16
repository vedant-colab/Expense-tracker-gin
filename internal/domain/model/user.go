package models

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      string    `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
