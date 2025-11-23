package user

import (
	dto "exptracker/internal/domain/dto/user"
	"exptracker/internal/middleware"
	svc "exptracker/internal/service/user"
	"exptracker/pkg/response"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service svc.Service
}

func NewUserController(s svc.Service) *Controller {
	return &Controller{s}
}

func (uc *Controller) GetProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)

	user, err := uc.service.GetProfile(userID)
	if err != nil {
		response.Error(c, 404, "USER_NOT_FOUND", err.Error())
		return
	}

	response.Success(c, 200, user)
}

func (uc *Controller) Update(c *gin.Context) {
	var req dto.UpdateUserRequest
	userID := middleware.GetUserID(c)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "INVALID_PAYLOAD", err.Error())
		return
	}

	user, err := uc.service.Update(userID, req)
	if err != nil {
		response.Error(c, 500, "UPDATE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, user)
}

func (uc *Controller) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	userID := middleware.GetUserID(c)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "INVALID_PAYLOAD", err.Error())
		return
	}

	if err := uc.service.ChangePassword(userID, req); err != nil {
		response.Error(c, 400, "PASSWORD_CHANGE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, gin.H{"message": "password updated"})
}

func (uc *Controller) Delete(c *gin.Context) {
	userID := middleware.GetUserID(c)

	if err := uc.service.Delete(userID); err != nil {
		response.Error(c, 500, "DELETE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, gin.H{"message": "account deleted"})
}
