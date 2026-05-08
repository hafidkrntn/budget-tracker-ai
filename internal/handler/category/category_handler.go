package category

import (
	"backend-go/internal/model/form"
	"backend-go/internal/service/category"
	"backend-go/pkg/response"
	"runtime"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var req form.Category
	_, file, line, _ := runtime.Caller(0)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetErrorFailedCreate(file, line)
		response.SetFailedCreateResponse(c, err)
		return
	}

	results, err := category.CreateCategory(req)
	if err != nil {
		response.SetErrorFailedCreate(file, line)
		response.SetFailedCreateResponse(c, err)
		return
	}

	response.SetSuccessCreateResponse(c, results)
}

func GetCategoryPagination(c *gin.Context) {
	_, file, line, _ := runtime.Caller(0)
	params := &form.Params{}
	if err := c.ShouldBindQuery(params); err != nil {
		response.SetFailedReadResponse(c, err)
		return
	}

	results, err := category.GetAllPagination(*params)
	if err != nil {
		response.SetErrorFailedRead(file, line)
		response.SetFailedReadResponse(c, err)
		return
	}

	response.SetSuccessReadResponse(c, results)
}

func GetCategory(c *gin.Context) {
	_, file, line, _ := runtime.Caller(0)

	results, err := category.GetAllCategory()
	if err != nil {
		response.SetErrorFailedRead(file, line)
		response.SetFailedReadResponse(c, err)
		return
	}

	response.SetSuccessReadResponse(c, results)
}

func GetCategoryById(c *gin.Context) {
	_, file, line, _ := runtime.Caller(0)

	id := c.Param("id")

	results, err := category.GetCategoryById(id)
	if err != nil {
		response.SetErrorFailedRead(file, line)
		response.SetFailedReadResponse(c, err)
		return
	}

	response.SetSuccessReadResponse(c, results)
}

func UpdateCategory(c *gin.Context) {
	var req form.Category
	_, file, line, _ := runtime.Caller(0)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetErrorFailedUpdate(file, line)
		response.SetFailedUpdateResponse(c, err)
		return
	}

	results, err := category.UpdateCategory(req)
	if err != nil {
		response.SetErrorFailedUpdate(file, line)
		response.SetFailedUpdateResponse(c, err)
		return
	}

	response.SetSuccessUpdateResponse(c, results)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	_, file, line, _ := runtime.Caller(0)

	results, err := category.DeletedCategory(id)
	if err != nil {
		response.SetErrorFailedDelete(file, line)
		response.SetFailedDeleteResponse(c, err)
		return
	}

	response.SetSuccessDeleteResponse(c, results)
}
