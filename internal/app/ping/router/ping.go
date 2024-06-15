package router

import (
	"DatabaseLab/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func AppPingInit(e *gin.Engine) {
	e.GET("/ping/v1", func(c *gin.Context) {
		response.HTTPSuccess(c, "pong")
	})

	e.GET("/ping/v1/err", func(c *gin.Context) {
		response.ServiceErr(c, errors.New("this is err"))
	})
}
