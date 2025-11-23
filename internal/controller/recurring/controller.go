package recurring

import (
	dto "exptracker/internal/domain/dto/recurring"
	"exptracker/internal/middleware"
	svc "exptracker/internal/service/recurring"
	"exptracker/pkg/response"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service svc.Service
}

func NewRecurringController(s svc.Service) *Controller {
	return &Controller{s}
}

func (rc *Controller) Create(c *gin.Context) {
	var req dto.CreateRecurringRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "INVALID_PAYLOAD", err.Error())
		return
	}

	userID := middleware.GetUserID(c)

	r, err := rc.service.Create(req, userID)
	if err != nil {
		response.Error(c, 400, "CREATE_FAILED", err.Error())
		return
	}

	response.Success(c, 201, r)
}

func (rc *Controller) GetAll(c *gin.Context) {
	userID := middleware.GetUserID(c)

	list, err := rc.service.GetAll(userID)
	if err != nil {
		response.Error(c, 500, "FETCH_FAILED", err.Error())
		return
	}

	response.Success(c, 200, list)
}

func (rc *Controller) Update(c *gin.Context) {
	var req dto.UpdateRecurringRequest
	userID := middleware.GetUserID(c)
	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "INVALID_PAYLOAD", err.Error())
		return
	}

	r, err := rc.service.Update(id, req, userID)
	if err != nil {
		response.Error(c, 400, "UPDATE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, r)
}

func (rc *Controller) Delete(c *gin.Context) {
	userID := middleware.GetUserID(c)
	id := c.Param("id")

	if err := rc.service.Delete(id, userID); err != nil {
		response.Error(c, 400, "DELETE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, gin.H{"message": "deleted"})
}
