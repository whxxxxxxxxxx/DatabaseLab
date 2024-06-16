package dao

import (
	"DatabaseLab/internal/app/users/model"
	"gorm.io/gorm"
)

type users struct {
	*gorm.DB
}

func (dao *users) Init(db *gorm.DB) (err error) {
	dao.DB = db
	return db.AutoMigrate(&model.Users{})
}

func (dao *users) ExistOrNotByUerName(userName string) (user *model.Users, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Users{}).Where("user_name = ?", userName).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, err
}

func (dao *users) CreateUser(user *model.Users) error {
	return dao.DB.Model(&model.Users{}).Create(user).Error
}

func (dao *users) GetUserById(id uint) (user *model.Users, err error) {
	err = dao.DB.Model(&model.Users{}).Where("id = ?", id).Find(&user).Error
	return
}
