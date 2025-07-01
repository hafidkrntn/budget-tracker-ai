package constants

var (
	ErrMissingLoginValues   = "missing Username or Password"
	ErrFailedAuthentication = "incorrect Username or Password"
	ErrExpiredToken         = "token is expired"
	ErrExistEmail           = "Email Already Registered"
	ErrConfirmPassword      = "Password is not the same"
	ErrEmptyAuthHeader      = "auth header is empty"
	ErrInvalidToken         = "Invalid Token"
	ErrLogout               = "Account Has Been Logged Out"
)

var (
	DataFound            = "Data Found"
	SuccessAddData       = "Successfully Added Data"
	SuccessUpdateData    = "Successfully Updated Data"
	SuccessDeleteData    = "Successfully Deleted Data"
	SuccessDownload      = "Successfully Download"
	SuccessDisplayedData = "Successfully Displayed The Data"
)

var (
	ErrDataNotFound     = "Data Not Found"
	ErrInternalServer   = "Internal Server Error"
	DataIsAvailable     = "Data Is Available"
	ContentNotFound     = "Content Not Found"
	ModuleNotFound      = "Module Not Found"
	FailedAddData       = "Failed To Added Data"
	FailedUpdateData    = "Failed To Change Data"
	FailedDeleteData    = "Failed deleted Data"
	FailedDownload      = "Failed Download"
	FailedDisplayedData = "Failed Displayed The Data"
)