package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
	"github.com/spf13/cobra"
)

var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

// appStartCommand 启动一个Web服务
var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个Web服务",
	RunE: func(c *cobra.Command, args []string) error {
		// 从服务容器中获取kernel的服务实例
		kernelService := framework.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		core := kernelService.WebEngine()

		server := &http.Server{
			Addr:    ":8080",
			Handler: core,
		}
		server.RegisterOnShutdown(
			func() {
				fmt.Println("the server is in shutdown")
			},
		)
		go func() {
			server.ListenAndServe()
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Server Shutdown:", err)
		}

		return nil
	},
}

func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}
