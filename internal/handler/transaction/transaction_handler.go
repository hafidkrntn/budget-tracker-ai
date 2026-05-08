package transaction

import (
	"backend-go/internal/model/form"
	"backend-go/internal/service/transaction"
	"backend-go/pkg/response"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getUserID(c *gin.Context) (uuid.UUID, bool) {
	raw, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, false
	}
	str, ok := raw.(string)
	if !ok {
		return uuid.Nil, false
	}
	id, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil, false
	}
	return id, true
}

func CreateTransaction(c *gin.Context) {
	var req form.TransactionForm
	_, file, line, _ := runtime.Caller(0)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetErrorFailedCreate(file, line)
		response.SetFailedCreateResponse(c, err)
		return
	}

	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing user identity in token"})
		return
	}
	req.UserId = userID

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
	search := c.DefaultQuery("search", "")

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

	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing user identity in token"})
		return
	}
	req.UserId = userID

	results, err := transaction.UpdateTransaction(req)
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
