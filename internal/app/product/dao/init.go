package dao

import (
	"gorm.io/gorm"
)

var (
	Product = &product{}
)

func InitOP(db *gorm.DB) error {
	err := Product.Init(db)
	if err != nil {
		return err
	}

	return err
}
