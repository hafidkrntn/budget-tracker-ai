package router

import (
	"backend-go/internal/handler/category"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.RouterGroup) {
	urlGroup := router.Group("category")
	urlGroup.POST("/insert", category.CreateCategory)
	urlGroup.GET("/pagination", category.GetCategoryPagination)
	urlGroup.GET("/get", category.GetCategory)
	urlGroup.GET("/get/:id", category.GetCategoryById)
	urlGroup.PUT("/update", category.UpdateCategory)
	urlGroup.DELETE("/delete/:id", category.DeleteCategory)
}
