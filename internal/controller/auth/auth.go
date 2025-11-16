package auth

import (
	dto "exptracker/internal/domain/dto/auth"
	"exptracker/internal/service/auth"
	"exptracker/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	authService auth.Service
}

func NewAuthController(authService auth.Service) *Controller {
	return &Controller{authService: authService}
}

func (ac *Controller) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_PAYLOAD", err.Error())
		return
	}

	ip := c.ClientIP()
	ua := c.GetHeader("User-Agent")

	resp, err := ac.authService.Register(req, ip, ua)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "REGISTER_FAILED", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, resp)
}

func (ac *Controller) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_PAYLOAD", err.Error())
		return
	}

	ip := c.ClientIP()
	ua := c.GetHeader("User-Agent")

	resp, err := ac.authService.Login(req, ip, ua)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", err.Error())
		return
	}

	response.Success(c, http.StatusOK, resp)
}

func (ac *Controller) Refresh(c *gin.Context) {
	var req dto.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_PAYLOAD", err.Error())
		return
	}

	ip := c.ClientIP()
	ua := c.GetHeader("User-Agent")

	resp, err := ac.authService.Refresh(req, ip, ua)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "REFRESH_FAILED", err.Error())
		return
	}

	response.Success(c, http.StatusOK, resp)
}
