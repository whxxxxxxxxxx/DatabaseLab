package errorx

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrorExistUser          = 30001
	ErrorFailEncrypt        = 30002
	ErrorUserNotFound       = 30003
	ErrorPasswordWrong      = 30004
	ErrorAuthToken          = 30005
	ErrorEmailFormat        = 30006
	ErrorImgUpload          = 30007
	ErrorAuthCheckTokenFail = 30008
)
