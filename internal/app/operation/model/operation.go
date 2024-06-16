package model

import (
	pmodel "DatabaseLab/internal/app/product/model"
	umodel "DatabaseLab/internal/app/users/model"

	"gorm.io/gorm"
)

// 数据库模型

type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint `gorm:"not null"`
	BossID    uint
	Num       uint
	MaxNum    uint
	Check     bool
}

type Favorite struct {
	gorm.Model
	User      umodel.Users   `gorm:"ForeignKey:UserID"`
	UserID    uint           `gorm:"not null"`
	Product   pmodel.Product `gorm:"ForeignKey:ProductID"`
	ProductID uint           `gorm:"not null"`
	Boss      umodel.Users   `gorm:"ForeignKey:BossID"`
	BossID    uint           `gorm:"not null"`
}

type Order struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	BossID    uint   `gorm:"not null"`
	AddressID uint   `gorm:"not null"`
	Num       int    // 数量
	OrderNum  uint64 // 订单号
	Type      uint   // 1 未支付  2 已支付
	Money     float64
}

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Address string `gorm:"type:varchar(50) not null"`
}
