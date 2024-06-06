package dao

import (
	"IOTProject/internal/app/ping/model"
	"IOTProject/store/mysql"
)

var (
	Ping *mysql.Orm
)

func AutoMigrate() error {
	return Ping.AutoMigrate(&model.Ping{})
}
