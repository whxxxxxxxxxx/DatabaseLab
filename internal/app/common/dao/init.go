package dao

import (
	"gorm.io/gorm"
)

var (
	Common = &common{}
)

func InitOP(db *gorm.DB) error {
	err := Common.Init(db)
	if err != nil {
		return err
	}

	return err
}
