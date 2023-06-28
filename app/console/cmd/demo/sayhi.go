package demo

import (
	"fmt"
	"time"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
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
	settingService := framework.MustMake(contract.SettingKey).(contract.Setting)
	distributedService := framework.MustMake(contract.DistributedKey).(contract.Distributed)
	localAppID := settingService.AppID()
	appID, err := distributedService.Select("sayhi_per_5s", localAppID, 2*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	if appID != localAppID {
		return
	}
	fmt.Println(settingService.AppID(), "-", time.Now(), "-Hello, Zima Framework is powerful")
}
