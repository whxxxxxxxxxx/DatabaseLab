package router

import (
	"DatabaseLab/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func AppProductInit(e *gin.Engine) {
	e.GET("/product/v1", func(c *gin.Context) {
		response.HttpResponse(c, response.ReturnResponse(200000, "users Init test", nil, nil))
	})

	e.GET("/product/v1/err", func(c *gin.Context) {
		response.ServiceError(c, errors.New("error test"))
	})
}
