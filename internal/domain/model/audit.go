package models

import "time"

type AuditLog struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	UserID    string    `gorm:"type:uuid;index" json:"user_id"`
	Action    string    `gorm:"type:varchar(255)" json:"action"`
	Entity    string    `gorm:"type:varchar(50)" json:"entity"`
	EntityID  string    `gorm:"type:uuid" json:"entity_id"`
	Meta      string    `gorm:"type:jsonb" json:"meta"`
	CreatedAt time.Time `json:"created_at"`
}
