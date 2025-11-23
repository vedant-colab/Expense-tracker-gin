package account

import (
	dto "exptracker/internal/domain/dto/account"
	"exptracker/internal/middleware"
	svc "exptracker/internal/service/account"
	"exptracker/pkg/response"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service svc.Service
}

func NewAccountController(s svc.Service) *Controller {
	return &Controller{s}
}

func (ac *Controller) Create(c *gin.Context) {
	var req dto.CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "INVALID_PAYLOAD", err.Error())
		return
	}

	userID := middleware.GetUserID(c)

	acc, err := ac.service.Create(req, userID)
	if err != nil {
		response.Error(c, 500, "CREATE_FAILED", err.Error())
		return
	}

	response.Success(c, 201, acc)
}

func (ac *Controller) GetAll(c *gin.Context) {
	userID := middleware.GetUserID(c)

	list, err := ac.service.GetAll(userID)
	if err != nil {
		response.Error(c, 500, "FETCH_FAILED", err.Error())
		return
	}

	response.Success(c, 200, list)
}

func (ac *Controller) GetByID(c *gin.Context) {
	userID := middleware.GetUserID(c)
	id := c.Param("id")

	acc, err := ac.service.GetByID(id, userID)
	if err != nil {
		response.Error(c, 404, "NOT_FOUND", err.Error())
		return
	}

	response.Success(c, 200, acc)
}

func (ac *Controller) Update(c *gin.Context) {
	var req dto.UpdateAccountRequest
	userID := middleware.GetUserID(c)
	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "INVALID_PAYLOAD", err.Error())
		return
	}

	acc, err := ac.service.Update(id, req, userID)
	if err != nil {
		response.Error(c, 400, "UPDATE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, acc)
}

func (ac *Controller) Delete(c *gin.Context) {
	userID := middleware.GetUserID(c)
	id := c.Param("id")

	if err := ac.service.Delete(id, userID); err != nil {
		response.Error(c, 400, "DELETE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, gin.H{"message": "deleted"})
}
