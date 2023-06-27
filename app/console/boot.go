package console

import (
	"github.com/lixiaoqing18/zima/app/console/cmd/demo"
	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/cmd"
	"github.com/spf13/cobra"
)

// RunCommand  初始化根 Command 并运行
func RunCommand() error {
	// 根 Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "zima",
		// 简短介绍
		Short: "zima 命令",
		// 根命令的详细介绍
		Long: "zima 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现 cobra 默认的 completion 子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	// 绑定框架的命令
	cmd.AddSysCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)
	//绑定cron命令
	AddCronCommand(rootCmd)
	// 执行 RootCommand
	return rootCmd.Execute()
}

func AddAppCommand(root *cobra.Command) {
	root.AddCommand(demo.InitFooCommand())

}

func AddCronCommand(root *cobra.Command) {
	root.AddCommand(demo.TimeCommand)
	framework.AddCron("* * * * * *", demo.TimeCommandFunc, demo.TimeCommand)

	root.AddCommand(demo.SayhiCommand)
	framework.AddCron("5 * * * * *", demo.SayhiCommandFunc, demo.SayhiCommand)
}
