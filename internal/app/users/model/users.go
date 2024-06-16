package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string `gorm:"not null"`
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          int
}
