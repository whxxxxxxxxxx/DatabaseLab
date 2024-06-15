package server

import (
	"DatabaseLab/config"
	"DatabaseLab/internal/app/appInitialize"
	"DatabaseLab/kernel"
	"DatabaseLab/pkg/ip"
	"DatabaseLab/pkg/stringx"
	"DatabaseLab/store/openGauss"
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	configYml string
	engine    *kernel.Engine
	StartCmd  = &cobra.Command{
		Use:     "server",
		Short:   "Set Application config info",
		Example: "main server -c config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			setUp()
			loadStore()
			loadApp()
		},
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/config.yaml", "Start server with provided configuration file")
}

// 初始化配置和日志
func setUp() {
	// 初始化全局 ctx
	ctx, cancel := context.WithCancel(context.Background())

	// 初始化资源管理器
	engine = &kernel.Engine{Ctx: ctx, Cancel: cancel}
	kernel.Kernel = engine

	// 加载配置
	config.LoadConfig(configYml, func(globalConfig *config.GlobalConfig) {
		for _, listener := range engine.ConfigListener {
			listener(globalConfig)
		}
	})

	//初始化Gin
	mode := config.GetConfig().MODE
	if mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 初始化 Gin
	engine.GIN = gin.Default()
	engine.GIN.Use(cors.Default())
}

// 存储介质连接
func loadStore() {
	engine.OpenGauss = openGauss.MustNewMysqlOrm(config.GetConfig().OpenGauss)
}

// 加载应用，包含多个生命周期
func loadApp() {
	apps := appInitialize.GetApps()
	for _, app := range apps {
		_err := app.PreInit(engine)
		if _err != nil {
			fmt.Errorf("app %s PreInit failed: %s", app.Info(), _err)
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.Init(engine)
		if _err != nil {
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.PostInit(engine)
		if _err != nil {
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.Load(engine)
		if _err != nil {
			os.Exit(1)
		}
	}
	for _, app := range apps {
		_err := app.Start(engine)
		if _err != nil {
			os.Exit(1)
		}
	}

}

// 启动服务
func run() {
	port := config.GetConfig().Port

	engine.HttpServer = &http.Server{
		Addr:    ":8001",
		Handler: engine.GIN,
	}

	go func() {
		// 服务连接
		if err := engine.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf(stringx.Green("listen: %s\n"), port)
			fmt.Printf(stringx.Yellow("Server run failed: %s\n"), err)
		}
	}()

	println(stringx.Green("Server run at:"))
	println(fmt.Sprintf("-  Local:   http://localhost:%s", port))
	localHost := ip.GetLocalHost()
	engine.CurrentIpList = make([]string, 0, len(localHost))
	for _, host := range localHost {
		engine.CurrentIpList = append(engine.CurrentIpList, host)
		println(fmt.Sprintf("-  Network: http://%s:%s", host, port))
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	println(stringx.Blue("Shutting down server..."))

	ctx, cancel := context.WithTimeout(engine.Ctx, 5*time.Second)
	defer engine.Cancel()
	defer cancel()

	if err := engine.HttpServer.Shutdown(ctx); err != nil {
		println(stringx.Yellow("Server forced to shutdown: " + err.Error()))
	}

	println(stringx.Green("Server exiting Correctly"))
}
