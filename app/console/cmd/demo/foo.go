package demo

import (
	"fmt"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/spf13/cobra"
)

var fooCommand = &cobra.Command{
	Use:     "foo",
	Short:   "foo的短说明",
	Long:    "foo的长说明",
	Aliases: []string{"fo", "f"},
	Example: "foo命令的例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(framework.GetContainer())
		return nil
	},
}

// appStartCommand 启动一个Web服务
var barCommand = &cobra.Command{
	Use:     "bar",
	Short:   "bar的短说明",
	Long:    "bar的长说明",
	Aliases: []string{"ba", "b"},
	Example: "bar命令的例子",
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println(framework.GetContainer())
		return nil
	},
}

func InitFooCommand() *cobra.Command {
	fooCommand.AddCommand(barCommand)
	return fooCommand
}
