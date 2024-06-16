package response

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Error   error       `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func ReturnResponse(code int, msg string, data interface{}, err error) JsonResponse {
	return JsonResponse{
		Code:    code,
		Message: msg,
		Data:    data,
		Error:   err,
	}
}

func HttpResponse(c *gin.Context, Res JsonResponse) {
	code := Res.Code
	msg := Res.Message
	data := Res.Data
	err := Res.Error
	if err != nil {
		c.JSON(int(code/1000), &ResponseData{
			Code:    code,
			Message: msg,
			Data:    data,
			Error:   err.Error(),
		})
	} else {
		c.JSON(int(code/1000), &ResponseData{
			Code:    code,
			Message: msg,
			Data:    data,
			Error:   "",
		})
	}

}

func ServiceError(c *gin.Context, err error) {
	HttpResponse(c, ReturnResponse(400, "内部异常", nil, err))
}
