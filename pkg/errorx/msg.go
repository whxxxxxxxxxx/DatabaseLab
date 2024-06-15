package errorx

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	ErrorExistUser:     "用户已存在",
	ErrorFailEncrypt:   "密码加密失败",
	ErrorUserNotFound:  "用户不存在",
	ErrorPasswordWrong: "密码错误",
	ErrorAuthToken:     "token 认证失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
