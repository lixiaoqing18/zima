package demo

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var TimeCommand = &cobra.Command{
	Use:     "printtime",
	Short:   "定时输出时间",
	Long:    "每隔一秒输出系统时间",
	Aliases: []string{"pt", "t"},
	Example: "printtime命令的例子",
	Run: func(cmd *cobra.Command, args []string) {
		TimeCommandFunc()
	},
}

func TimeCommandFunc() {
	fmt.Println(time.Now())
}
