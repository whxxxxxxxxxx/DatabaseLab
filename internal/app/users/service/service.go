package service

import (
	"DatabaseLab/internal/app/users/dao"
	"DatabaseLab/internal/app/users/dto"
	"DatabaseLab/internal/app/users/model"
	"DatabaseLab/middleware/response"
	"DatabaseLab/pkg/ctl"
	"DatabaseLab/pkg/emailx"
	"DatabaseLab/pkg/errorx"
	"DatabaseLab/pkg/jwt"
	"DatabaseLab/pkg/passwdx"
	"errors"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, service dto.UserRegisterService) response.JsonResponse {
	code := errorx.Success
	userDao := dao.Users
	_, exist, err := userDao.ExistOrNotByUerName(service.UserName)
	if err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	if exist {
		code = errorx.ErrorExistUser
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	//正则检查email格式
	if service.Email != "" && !emailx.CheckEmail(service.Email) {
		code = errorx.ErrorEmailFormat
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	user := model.Users{
		NickName: service.NickName,
		UserName: service.UserName,
		Avatar:   service.Avatar,
		Money:    0,
		Email:    service.Email,
		Status:   passwdx.Active,
	}
	err = passwdx.SetPassword(service.Password, &user)
	if err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	err = userDao.CreateUser(&user)
	if err != nil {
		code = errorx.Error
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
}

func Login(ctx *gin.Context, service dto.UserLoginService) response.JsonResponse {
	code := errorx.Success
	userDao := dao.Users
	user, exist, err := userDao.ExistOrNotByUerName(service.UserName)
	if err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	if !exist {
		code = errorx.ErrorUserNotFound
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, errors.New("用户不存在，请先注册"))
	}
	if !passwdx.CheckPassword(service.Password, user) {
		code = errorx.ErrorPasswordWrong
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}

	token, refreshToken, err := jwt.GenerateToken(user.ID, user.NickName)
	if err != nil {
		code = errorx.ErrorAuthToken
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), dto.TokenData{Token: token, RefreshToken: refreshToken,
		User: dto.BuildUser(user)}, nil)

}

func UserInfo(ctx *gin.Context) response.JsonResponse {
	u, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		return response.ReturnResponse(errorx.Error, errorx.GetMsg(errorx.Error), nil, err)
	}
	code := errorx.Success
	userDao := dao.Users
	user, err := userDao.GetUserById(u.Id)
	if err != nil {
		code = errorx.Error
		return response.ReturnResponse(code, errorx.GetMsg(code), nil, nil)
	}
	return response.ReturnResponse(code, errorx.GetMsg(code), dto.BuildUser(user), nil)
}
