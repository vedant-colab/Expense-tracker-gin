package response

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func Error(c *gin.Context, httpCode int, code string, msg string) {
	c.JSON(httpCode, ErrorResponse{
		Status:  "error",
		Message: msg,
		Code:    code,
	})
}
