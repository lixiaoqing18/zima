package cmd

import (
	"fmt"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/spf13/cobra"
)

var cronCommand = &cobra.Command{
	Use:   "cron",
	Short: "cron任务控制命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

// appStartCommand 启动一个Web服务
var cronStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个cron任务",
	RunE: func(c *cobra.Command, args []string) error {
		framework.StartCron()
		fmt.Println("cron服务已启动")
		return nil
	},
}

func initCronCommand() *cobra.Command {
	cronCommand.AddCommand(cronStartCommand)
	return cronCommand
}
