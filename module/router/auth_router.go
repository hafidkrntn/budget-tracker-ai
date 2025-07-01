package router

import (
	"backend-go/module/handler/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	urlGroup := router.Group("/auth")
	urlGroup.POST("/register", auth.RegisterUsers)
	urlGroup.POST("/login", auth.LoginUsers)
}
