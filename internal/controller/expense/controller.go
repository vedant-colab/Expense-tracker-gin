package expense

import (
	"net/http"

	dto "exptracker/internal/domain/dto/expense"
	"exptracker/internal/middleware"
	service "exptracker/internal/service/expense"
	"exptracker/pkg/response"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.Service
}

func NewExpenseController(service service.Service) *Controller {
	return &Controller{service}
}

func (ec *Controller) Create(c *gin.Context) {
	var req dto.CreateExpenseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_PAYLOAD", err.Error())
		return
	}

	userID := middleware.GetUserID(c)

	exp, err := ec.service.Create(req, userID)
	if err != nil {
		response.Error(c, 500, "CREATE_FAILED", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, exp)
}

func (ec *Controller) GetAll(c *gin.Context) {
	userID := middleware.GetUserID(c)

	list, err := ec.service.GetAll(userID)
	if err != nil {
		response.Error(c, 500, "FETCH_FAILED", err.Error())
		return
	}

	response.Success(c, 200, list)
}

func (ec *Controller) GetByID(c *gin.Context) {
	userID := middleware.GetUserID(c)
	id := c.Param("id")

	exp, err := ec.service.GetByID(id, userID)
	if err != nil {
		response.Error(c, 404, "NOT_FOUND", err.Error())
		return
	}

	response.Success(c, 200, exp)
}

func (ec *Controller) Update(c *gin.Context) {
	var req dto.UpdateExpenseRequest
	userID := middleware.GetUserID(c)
	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "INVALID_PAYLOAD", err.Error())
		return
	}

	exp, err := ec.service.Update(id, req, userID)
	if err != nil {
		response.Error(c, 400, "UPDATE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, exp)
}

func (ec *Controller) Delete(c *gin.Context) {
	userID := middleware.GetUserID(c)
	id := c.Param("id")

	if err := ec.service.Delete(id, userID); err != nil {
		response.Error(c, 400, "DELETE_FAILED", err.Error())
		return
	}

	response.Success(c, 200, gin.H{"message": "deleted"})
}
