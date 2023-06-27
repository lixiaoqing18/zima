package demo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SayhiCommand = &cobra.Command{
	Use:     "sayhi",
	Short:   "定时输出Hello",
	Long:    "每隔5秒输出Hello",
	Aliases: []string{"hello", "hi"},
	Example: "sayhi命令的例子",
	Run: func(cmd *cobra.Command, args []string) {
		SayhiCommandFunc()
	},
}

func SayhiCommandFunc() {
	fmt.Println("Hello, Zima Framework is powerful")
}
