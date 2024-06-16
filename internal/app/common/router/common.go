package router

import (
	"DatabaseLab/internal/app/common/handler"
	"DatabaseLab/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func AppCommonInit(e *gin.Engine) {
	e.GET("/common/v1", func(c *gin.Context) {
		response.HttpResponse(c, response.ReturnResponse(200000, "users Init test", nil, nil))
	})

	e.GET("/common/v1/err", func(c *gin.Context) {
		response.ServiceError(c, errors.New("error test"))
	})

	e.POST("/common/upload", handler.HandleUpload)
}
