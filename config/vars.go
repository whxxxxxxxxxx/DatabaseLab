package config

import (
	"DatabaseLab/store/openGauss"
)

type GlobalConfig struct {
	MODE        string            `yaml:"Mode"`
	ProgramName string            `yaml:"ProgramName"`
	BaseURL     string            `yaml:"BaseURL"`
	AUTHOR      string            `yaml:"Author"`
	Listen      string            `yaml:"Listen"`
	Port        string            `yaml:"Port"`
	OpenGauss   openGauss.OrmConf `yaml:"OpenGauss"`
}
