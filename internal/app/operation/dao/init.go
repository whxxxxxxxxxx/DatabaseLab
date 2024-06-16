package dao

import (
	"gorm.io/gorm"
)

var (
	Operation = &operation{}
)

func InitOP(db *gorm.DB) error {
	err := Operation.Init(db)
	if err != nil {
		return err
	}

	return err
}
