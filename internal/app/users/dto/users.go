package dto

import "DatabaseLab/internal/app/users/model"

type User struct {
	ID       uint   `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

type UserRegisterService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

type UserLoginService struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

func BuildUser(user *model.Users) *User {
	return &User{
		ID:       user.ID,
		NickName: user.NickName,
		UserName: user.UserName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
