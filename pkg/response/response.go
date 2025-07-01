package response

import (
	"backend-go/constants"
	logerror "backend-go/module/repository/log_error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

func SetResponseJSON(c *gin.Context, code int, data interface{}, msg string, err error) {
	response := Response{
		Code:  code,
		Data:  data,
		Msg:   msg,
		Error: "",
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(code, response)
	c.Abort()
}

var LogRepo *logerror.LogErrorRepository

func logError(message string, file string, code int, line int) {
	if LogRepo != nil {
		go LogRepo.CreateLogError(message, file, code, line)
	}
}

func SetErrorFailedCreate(file string, line int) {
	logError(constants.FailedAddData, file, http.StatusBadRequest, line)
}

func SetErrorFailedRead(file string, line int) {
	logError(constants.FailedDisplayedData, file, http.StatusBadRequest, line)
}

func SetErrorEmptyID(file string, line int) {
	logError(constants.DataFound, file, http.StatusOK, line)
}

func SetErrorFailedUpdate(file string, line int) {
	logError(constants.FailedUpdateData, file, http.StatusBadRequest, line)
}

func SetErrorFailedUpdateHeader(file string, line int) {
	logError(constants.FailedUpdateData+" Vendor ID is Required", file, http.StatusBadRequest, line)
}

func SetErrorFailedDelete(file string, line int) {
	logError(constants.FailedDeleteData, file, http.StatusBadRequest, line)
}

func SetErrorStatusNotFound(file string, line int) {
	logError(constants.ErrLogout, file, http.StatusNotFound, line)
}

func SetErrorInvalidToken(file string, line int) {
	logError(constants.ErrInvalidToken, file, http.StatusUnauthorized, line)
}

func SetErrorDataNotFound(file string, line int) {
	logError(constants.ErrDataNotFound, file, http.StatusBadRequest, line)
}

func SetErrorEmptyCurrency(file string, line int) {
	logError(constants.SuccessDisplayedData, file, http.StatusOK, line)
}

func SetErrorInternalServer(file string, line int) {
	logError(constants.ErrInternalServer, file, http.StatusInternalServerError, line)
}

func SetErrorEmptyHeader(file string, line int) {
	logError(constants.ErrEmptyAuthHeader, file, http.StatusUnauthorized, line)
}

func SetErrorContentNotMultipart(file string, line int) {
	logError("Content-Type isn't multipart/form-data", file, http.StatusBadRequest, line)
}

func SetErrorRetrievingFile(file string, line int) {
	logError("Error retrieving the file", file, http.StatusBadRequest, line)
}

func SetErrorModuleRequired(file string, line int) {
	logError("Module is required", file, http.StatusBadRequest, line)
}

func SetErrorUploadingFile(file string, line int) {
	logError("Error uploading the file", file, http.StatusInternalServerError, line)
}

func SetErrLogoutResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrLogout, err)
}

func SetUnauthorizedResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrInvalidToken, err)
}

func SetEmptyCurrencyResponse(c *gin.Context, data any, err error) {
	SetResponseJSON(c, http.StatusOK, gin.H{}, constants.SuccessDisplayedData, err)
}

func SetEmptyCompanyNameResponse(c *gin.Context, data any, err error) {
	SetResponseJSON(c, http.StatusOK, gin.H{}, constants.SuccessDisplayedData, err)
}

func SetFailedDeleteDataResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDeleteData, err)
}

func SetSuccessDeleteDataResponse(c *gin.Context, data any) {
	SetResponseJSON(c, http.StatusOK, data, constants.SuccessDeleteData, nil)
}

func SetStatusForbiddenResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusForbidden, nil, constants.ErrFailedAuthentication, nil)
}

func SetUnauthorizedEmptyAuthHeader(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusUnauthorized, nil, constants.ErrEmptyAuthHeader, nil)
}

func SetDataNotFoundResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusBadRequest, nil, constants.ErrDataNotFound, err)
}

func SetFailedCreateResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
}

func SetSuccessCreateResponse(c *gin.Context, data any) {
	SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, nil)
}

func SetFailedReadResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
}

func SetSuccessReadResponse(c *gin.Context, data any) {
	SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, nil)
}

func SetFailedUpdateResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
}

func SetSuccessUpdateResponse(c *gin.Context, data any) {
	SetResponseJSON(c, http.StatusOK, data, constants.SuccessUpdateData, nil)
}

func SetFailedDeleteResponse(c *gin.Context, err error) {
	SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDeleteData, err)
}

func SetSuccessDeleteResponse(c *gin.Context, data any) {
	SetResponseJSON(c, http.StatusOK, data, constants.SuccessDeleteData, nil)
}
