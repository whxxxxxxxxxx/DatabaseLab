package dao

import (
	"DatabaseLab/internal/app/operation/model"
	"gorm.io/gorm"
)

type operation struct {
	*gorm.DB
}

func (u *operation) Init(db *gorm.DB) (err error) {
	u.DB = db
	return db.AutoMigrate(&model.Cart{}, &model.Favorite{}, &model.Order{}, &model.Address{})
}
