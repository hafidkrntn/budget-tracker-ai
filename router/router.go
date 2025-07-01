package router

import (
	"backend-go/middleware"
	"backend-go/module/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "🎉 Hello from Go! App is running successfully!",
			})
		})

		/* API */
		router.AuthRoutes(api)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())

		router.CategoryRoutes(protected)
		router.TransactionRoutes(protected)
	}

	return r
}
