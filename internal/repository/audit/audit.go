package audit

import (
	"exptracker/internal/database"
	models "exptracker/internal/domain/model"
)

type Repository interface {
	Create(log *models.AuditLog) error
}

type auditRepository struct{}

func NewAuditRepository() Repository {
	return &auditRepository{}
}

func (r *auditRepository) Create(log *models.AuditLog) error {
	return database.DB.Create(log).Error
}
