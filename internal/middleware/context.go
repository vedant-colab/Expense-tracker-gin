package middleware

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) string {
	userID, exists := c.Get("user_id")
	if !exists {
		return ""
	}
	return userID.(string)
}

func GetUserRole(c *gin.Context) string {
	role, _ := c.Get("user_role")
	return role.(string)
}
