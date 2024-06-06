package config

import (
	"IOTProject/store/mysql"
	"IOTProject/store/openGauss"
)

type GlobalConfig struct {
	MODE        string            `yaml:"Mode"`
	ProgramName string            `yaml:"ProgramName"`
	BaseURL     string            `yaml:"BaseURL"`
	AUTHOR      string            `yaml:"Author"`
	Listen      string            `yaml:"Listen"`
	Port        string            `yaml:"Port"`
	SKLMysql    mysql.OrmConf     `yaml:"SKLMysql"`
	OpenGauss   openGauss.OrmConf `yaml:"OpenGauss"`
}
