package dao

import (
	"gorm.io/gorm"
)

var (
	Users = &users{}
)

func InitOp(db *gorm.DB) error {
	err := Users.Init(db)
	if err != nil {
		return err
	}

	return err
}
