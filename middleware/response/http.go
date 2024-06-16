package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	Code    int    `json:"code"`
	Error   any    `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func ReturnResponse(code int, msg string, data any, err any) JsonResponse {
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
	if code == 0 {
		c.JSON(http.StatusOK, &JsonResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		})
		return
	}
	c.JSON(int(code/1000), &JsonResponse{
		Code:    code,
		Message: msg,
		Data:    data,
		Error:   err,
	})
}

func ServiceError(c *gin.Context, err ...any) {
	HttpResponse(c, ReturnResponse(400, "内部异常", nil, err))
}
