package handler

import (
	"DatabaseLab/internal/app/users/dto"
	"DatabaseLab/internal/app/users/service"
	"DatabaseLab/middleware/response"
	"github.com/gin-gonic/gin"
)

func HandleUserRegister(c *gin.Context) {
	var userRegister dto.UserRegisterService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := service.Register(c, userRegister)
		response.HttpResponse(c, res)
	} else {
		response.ServiceError(c, err)
	}
}

func HandleUserLogin(c *gin.Context) {
	var userLogin dto.UserLoginService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := service.Login(c, userLogin)
		response.HttpResponse(c, res)
	} else {
		response.ServiceError(c, err)
	}
}

func HandleUserInfo(c *gin.Context) {
	res := service.UserInfo(c)
	response.HttpResponse(c, res)
}
