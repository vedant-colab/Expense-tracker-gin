package middleware

import (
	"exptracker/internal/config"
	"exptracker/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWT(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing Authorization Header")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid Authorization header")
			c.Abort()
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(cfg.App.JWTSecret), nil
		})

		if err != nil || token.Valid {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid or expired token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid token claims")
			c.Abort()
			return
		}
		userID, ok := claims["user_id"].(string)
		if !ok || userID == "" {
			response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing user_id")
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			response.Error(c, 401, "INVALID_TOKEN", "role missing in token")
			c.Abort()
			return
		}
		c.Set("user_id", userID)
		c.Set("role", role)

		c.Next()
	}

}
