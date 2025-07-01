package middleware

import (
	"backend-go/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := token.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: " + err.Error(),
			})
			c.Abort()
			return
		}

		// Simpan claims ke context (biar bisa diakses di handler)
		c.Set("email", claims["email"])

		c.Next()
	}
}
