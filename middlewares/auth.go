package middlewares

import (
	"E_commerce_System/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Missing token"})
			return
		}

		// 拆分 Bearer token
		var tokenString string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			return
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid token format"})
			return
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid token", "error": err.Error()})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
