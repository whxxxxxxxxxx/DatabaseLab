package dao

import (
	"DatabaseLab/internal/app/ping/model"
	"DatabaseLab/store/openGauss"
)

var (
	Ping *openGauss.Orm
)

func AutoMigrate() error {
	return Ping.AutoMigrate(&model.Ping{})
}
