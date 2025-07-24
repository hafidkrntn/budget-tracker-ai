package router

import (
	"backend-go/internal/router"
	"backend-go/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Sesuaikan dengan asal request FE kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
