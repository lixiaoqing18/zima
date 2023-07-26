package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/lixiaoqing18/zima/app/provider/demo"
	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type DemoAPI struct {
}

func NewDemoAPI() *DemoAPI {
	return &DemoAPI{}
}

func Register(core *gin.Engine) error {
	err := framework.Bind(demo.NewServiceProvider())
	if err != nil {
		return err
	}
	api := NewDemoAPI()
	core.GET("/demo", api.demo1)
	return nil
}

func (api *DemoAPI) demo1(ctx *gin.Context) {
	settingService := framework.MustMake(contract.SettingKey).(contract.Setting)
	demoService := framework.MustMake(demo.Key).(demo.IService)

	ctx.JSON(200, gin.H{
		"foo":              demoService.Demo(),
		"BaseFolder":       settingService.BaseFolder(),
		"Version":          settingService.Version(),
		"ConfigFolder":     settingService.ConfigFolder(),
		"LogFolder":        settingService.LogFolder(),
		"ProviderFolder":   settingService.ProviderFolder(),
		"MiddlewareFolder": settingService.MiddlewareFolder(),
		"CommandFolder":    settingService.CommandFolder(),
		"RuntimeFolder":    settingService.RuntimeFolder(),
		"TestFolder":       settingService.TestFolder(),
	})

	framework.GetLog().Info(ctx, "demo1 response completed", map[string]any{})
}
