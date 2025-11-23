package middleware

import (
	"exptracker/pkg/response"
	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := GetUserRole(c)

        if userRole != role {
            response.Error(c, 403, "FORBIDDEN", "insufficient permissions")
		    c.Abort()
		    return
        }

		c.Next()
	}
}

func RequireAnyRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := GetUserRole(c)

		for _, r := range roles {
			if r == userRole {
				c.Next()
				return
			}
		}

		response.Error(c, 403, "FORBIDDEN", "insufficient permissions")
		c.Abort()
	}
}
