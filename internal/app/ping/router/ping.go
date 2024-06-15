package router

import (
	"DatabaseLab/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func AppPingInit(e *gin.Engine) {
	e.GET("/ping/v1", func(c *gin.Context) {
		response.HttpResponse(c, response.ReturnResponse(c, 200000, "ping Init test", nil, nil))
	})

	e.GET("/ping/v1/err", func(c *gin.Context) {
		response.ServiceError(c, errors.New("error test"))
	})
}
