package dto

import "DatabaseLab/internal/app/users/model"

type User struct {
	ID       uint   `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	Money    int    `json:"money"`
	CreateAt int64  `json:"create_at"`
}

type UserRegisterService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type UserLoginService struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

type UserUpdateService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	Email    string `json:"email" form:"email"`
	Avatar   string `json:"avatar" form:"avatar"`
}

func BuildUser(user *model.Users) *User {
	return &User{
		ID:       user.ID,
		NickName: user.NickName,
		UserName: user.UserName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.Avatar,
		Money:    user.Money,
		CreateAt: user.CreatedAt.Unix(),
	}
}
