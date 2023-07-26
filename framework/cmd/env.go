package cmd

import (
	"fmt"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
	"github.com/spf13/cobra"
)

var envCommand = &cobra.Command{
	Use:   "env",
	Short: "环境变量查看命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

var envListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出所有环境变量",
	RunE: func(c *cobra.Command, args []string) error {
		envService := framework.MustMake(contract.EnvKey).(contract.Env)
		for k, v := range envService.All() {
			fmt.Println(k, "=", v)
		}
		return nil
	},
}

func initEnvCommand() *cobra.Command {
	envCommand.AddCommand(envListCommand)
	return envCommand
}
