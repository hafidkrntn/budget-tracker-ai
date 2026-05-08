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

		userID, ok := claims["user_id"].(string)
		if !ok || userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: token does not contain user identity, please re-login",
			})
			c.Abort()
			return
		}

		// Simpan claims ke context (biar bisa diakses di handler)
		c.Set("email", claims["email"])
		c.Set("user_id", userID)

		c.Next()
	}
}
