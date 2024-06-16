package router

import (
	"DatabaseLab/internal/app/users/handler"
	"DatabaseLab/middleware"
	"DatabaseLab/middleware/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func AppUsersInit(e *gin.Engine) {
	e.GET("/users/v1", func(c *gin.Context) {
		response.HttpResponse(c, response.ReturnResponse(200000, "users Init test", nil, nil))
	})

	e.GET("/users/v1/err", func(c *gin.Context) {
		response.ServiceError(c, errors.New("error test"))
	})

	userRoute := e.Group("/users")
	{
		userRoute.POST("/register", handler.HandleUserRegister)
		userRoute.POST("/login", handler.HandleUserLogin)
	}

	authed := userRoute.Group("/")
	authed.Use(middleware.AuthMiddleware())
	{
		authed.GET("info", handler.HandleUserInfo)
	}
}
