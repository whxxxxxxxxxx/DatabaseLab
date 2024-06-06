package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	Code    int32  `json:"code"`
	Error   any    `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func httpResponse(c *gin.Context, code int32, msg string, data any, err any) {
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

func HTTPSuccess(c *gin.Context, data any) {
	httpResponse(c, 0, "success", data, nil)
}

func HTTPFail(c *gin.Context, code int, msg string, err ...any) {
	newErrs := make([]string, 0, len(err))
	for _, e := range err {
		if ve, ok := e.(error); ok {
			newErrs = append(newErrs, ve.Error())
			continue
		}
		if ve, ok := e.(string); ok {
			newErrs = append(newErrs, ve)
			continue
		}
	}
	httpResponse(c, int32(code), msg, nil, newErrs)
}

func HTTPFailWithData(c *gin.Context, code int, msg string, data any, err ...any) {
	for i, e := range err {
		if v, ok := e.(error); ok {
			err[i] = v.Error()
		}
	}
	httpResponse(c, int32(code), msg, data, err)
}

func UnAuthorization(c *gin.Context) {
	HTTPFail(c, 401, "登录过期失效，请重新登录")
}

func Forbidden(c *gin.Context) {
	HTTPFail(c, 403, "未授权")
}

func InValidParam(c *gin.Context, err ...any) {
	HTTPFail(c, 400, "请求校验失败", err...)
}

func ServiceErr(c *gin.Context, err ...any) {
	HTTPFail(c, 500, "内部异常", err...)
}
