package middleware

import (
	"net/http"

	"github.com/Rai-Sahil/backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("auth_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		uid, err := utils.ValidateJWT(cookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		
		c.Set("userId", uid)
		c.Next()
	}
}
