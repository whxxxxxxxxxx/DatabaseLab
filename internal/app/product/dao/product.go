package dao

import (
	"DatabaseLab/internal/app/product/model"
	"gorm.io/gorm"
)

type product struct {
	*gorm.DB
}

func (u *product) Init(db *gorm.DB) (err error) {
	u.DB = db
	return db.AutoMigrate(&model.ProductImage{}, &model.Product{}, &model.ProductInfoImg{}, &model.ProductParamsImg{}, &model.Category{}, &model.Carousel{})
}
