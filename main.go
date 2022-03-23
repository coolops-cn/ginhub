package main

import (
	"fmt"
	"os"

	"github.com/coolops-cn/ginhub/bootstrap"
	"github.com/coolops-cn/ginhub/pkg/config"
	"github.com/coolops-cn/ginhub/pkg/console"
	"github.com/spf13/cobra"

	"github.com/coolops-cn/ginhub/internal/app/cmd"
	internalConfig "github.com/coolops-cn/ginhub/internal/config"
)

func init() {
	internalConfig.InitConfig()
}

func main() {

	// 应用入口
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple project ginhub",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,
		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDatabase()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.ServeCommand,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.ServeCommand)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
