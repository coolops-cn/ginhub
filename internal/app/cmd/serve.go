package cmd

import (
	"net/http"
	"syscall"

	"github.com/coolops-cn/ginhub/bootstrap"
	"github.com/coolops-cn/ginhub/pkg/config"
	"github.com/coolops-cn/ginhub/pkg/shutdown"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// 运行服务

var ServeCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start WEB Server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

// runWeb 运行服务
func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	// 初始化 Gin 实例
	router := gin.New()

	// 初始化路由
	bootstrap.SetupRouter(router)

	server := &http.Server{
		Addr:    config.Get("app.port"),
		Handler: router,
	}

	// 运行服务
	go server.ListenAndServe()

	// 优雅退出
	quit := shutdown.New(10)
	quit.Add(syscall.SIGINT, syscall.SIGTERM)
	quit.Start(server)
}
