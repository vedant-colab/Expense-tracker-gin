package response

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, SuccessResponse{
		Status: "success",
		Data:   data,
	})
}
