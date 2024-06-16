package model

import (
	"gorm.io/gorm"
)

// 数据库模型

type ProductImage struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ProductId uint   `json:"product_id"`
	BossId    uint   `json:"boss_id"`
	ImgPath   string `json:"img_path"`
}

type ProductInfoImg struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ProductId uint   `json:"product_id"`
	BossId    uint   `json:"boss_id"`
	ImgPath   string `json:"img_path"`
}

type ProductParamsImg struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ProductId uint   `json:"product_id"`
	BossId    uint   `json:"boss_id"`
	ImgPath   string `json:"img_path"`
}

type Product struct {
	gorm.Model
	ProductName   string `json:"product_name"`
	CategoryId    uint   `json:"category_id" gorm:"not null"`
	Title         string `json:"title"`
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
}

type Category struct {
	gorm.Model
	CategoryName string `json:"category_name"`
}

type Carousel struct {
	gorm.Model
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
}
