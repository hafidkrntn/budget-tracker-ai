package auth

import (
	"backend-go/module/model/form"
	"backend-go/module/service/auth"
	"backend-go/pkg/response"
	"runtime"

	"github.com/gin-gonic/gin"
)

func RegisterUsers(c *gin.Context) {
	var req form.RegisterForm
	_, file, line, _ := runtime.Caller(0)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetErrorFailedCreate(file, line)
		response.SetFailedCreateResponse(c, err)
		return
	}

	data, err := auth.Register(req)
	if err != nil {
		response.SetErrorFailedCreate(file, line)
		response.SetFailedCreateResponse(c, err)
		return
	}

	response.SetSuccessCreateResponse(c, data)
}

func LoginUsers(c *gin.Context) {
	var req form.LoginForm
	_, file, line, _ := runtime.Caller(0)

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetErrorFailedCreate(file, line)
		response.SetFailedReadResponse(c, err)
		return
	}

	token, err := auth.Login(req)
	if err != nil || token.Token == "" {
		response.SetErrorInvalidToken(file, line)
		response.SetFailedReadResponse(c, err)
		return
	}

	response.SetSuccessReadResponse(c, token)
}
