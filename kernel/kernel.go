package kernel

import (
	"IOTProject/config"
	"IOTProject/store/mysql"
	"IOTProject/store/tdengine"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Engine struct {
		GIN *gin.Engine

		SKLMySQL  *mysql.Orm
		OpenGauss *tdengine.Orm

		Ctx    context.Context
		Cancel context.CancelFunc

		HttpServer *http.Server

		CurrentIpList []string

		ConfigListener []func(*config.GlobalConfig)
	}
)

var Kernel *Engine
