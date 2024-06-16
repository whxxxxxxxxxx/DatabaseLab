package dao

import (
	"DatabaseLab/internal/app/common/model"
	"gorm.io/gorm"
)

type common struct {
	*gorm.DB
}

func (u *common) Init(db *gorm.DB) (err error) {
	u.DB = db
	return db.AutoMigrate(&model.Common{})
}
