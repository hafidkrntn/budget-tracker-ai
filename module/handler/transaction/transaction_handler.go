package transaction

import (
	"backend-go/module/model/form"
	"backend-go/module/service/transaction"
	"backend-go/pkg/response"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var req form.TransactionForm
	_, file, line, _ := runtime.Caller(0)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetErrorFailedCreate(file, line)
		response.SetFailedCreateResponse(c, err)
		return
	}

	results, err := transaction.CreateTransaction(req)
	if err != nil {
		response.SetErrorFailedCreate(file, line)
		response.SetFailedCreateResponse(c, err)
		return
	}

	response.SetSuccessCreateResponse(c, results)
}

func GetTransactionPagination(c *gin.Context) {
	_, file, line, _ := runtime.Caller(0)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "none")

	params := form.TransactionParams{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

	results, err := transaction.GetTransactionPagination(params)
	if err != nil {
		response.SetErrorFailedRead(file, line)
		response.SetFailedReadResponse(c, err)
		return
	}

	response.SetSuccessReadResponse(c, results)
}

func GetTransactionById(c *gin.Context) {
	_, file, line, _ := runtime.Caller(0)

	id := c.Param("id")

	results, err := transaction.GetTransactionById(id)
	if err != nil {
		response.SetErrorFailedRead(file, line)
		response.SetFailedReadResponse(c, err)
		return
	}

	response.SetSuccessReadResponse(c, results)
}

func UpdateTransaction(c *gin.Context) {
	var req form.TransactionForm
	_, file, line, _ := runtime.Caller(0)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetErrorFailedUpdate(file, line)
		response.SetFailedUpdateResponse(c, err)
		return
	}

	results, err := transaction.CreateTransaction(req)
	if err != nil {
		response.SetErrorFailedUpdate(file, line)
		response.SetFailedUpdateResponse(c, err)
		return
	}

	response.SetSuccessUpdateResponse(c, results)
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	_, file, line, _ := runtime.Caller(0)

	results, err := transaction.DeleteTransaction(id)
	if err != nil {
		response.SetErrorFailedDelete(file, line)
		response.SetFailedDeleteResponse(c, err)
		return
	}

	response.SetSuccessDeleteResponse(c, results)
}
