package kernel

import (
	"DatabaseLab/config"
	"DatabaseLab/store/openGauss"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Engine struct {
		GIN *gin.Engine

		OpenGauss *openGauss.Orm

		Ctx    context.Context
		Cancel context.CancelFunc

		HttpServer *http.Server

		CurrentIpList []string

		ConfigListener []func(*config.GlobalConfig)
	}
)

var Kernel *Engine
