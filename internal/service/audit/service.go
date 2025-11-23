package audit

import (
	"encoding/json"
	"time"

	models "exptracker/internal/domain/model"
	"exptracker/internal/repository/audit"
)

type Service interface {
	Log(userID, action, entity, entityID string, meta interface{}) error
}

type auditService struct {
	repo audit.Repository
}

func NewAuditService(repo audit.Repository) Service {
	return &auditService{repo}
}

func (s *auditService) Log(userID, action, entity, entityID string, meta interface{}) error {

	metaJSON := ""
	if meta != nil {
		b, _ := json.Marshal(meta)
		metaJSON = string(b)
	}

	log := &models.AuditLog{
		UserID:    userID,
		Action:    action,
		Entity:    entity,
		EntityID:  entityID,
		Meta:      metaJSON,
		CreatedAt: time.Now(),
	}

	return s.repo.Create(log)
}
