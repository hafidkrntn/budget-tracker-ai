package router

import (
	"backend-go/internal/handler/transaction"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.RouterGroup) {
	urlGroup := router.Group("transaction")
	urlGroup.POST("/create", transaction.CreateTransaction)
	urlGroup.GET("/paginated", transaction.GetTransactionPagination)
	urlGroup.GET("/get/:id", transaction.GetTransactionById)
	urlGroup.PUT("/update", transaction.UpdateTransaction)
	urlGroup.DELETE("/delete/:id", transaction.DeleteTransaction)
}
