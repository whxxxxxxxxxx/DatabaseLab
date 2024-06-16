package router

import (
	"DatabaseLab/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func AppOperationInit(e *gin.Engine) {
	e.GET("/operation/v1", func(c *gin.Context) {
		response.HttpResponse(c, response.ReturnResponse(200000, "users Init test", nil, nil))
	})

	e.GET("/operation/v1/err", func(c *gin.Context) {
		response.ServiceError(c, errors.New("error test"))
	})
}
