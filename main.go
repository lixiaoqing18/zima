package main

import (
	"github.com/lixiaoqing18/zima/app/console"
	"github.com/lixiaoqing18/zima/app/web"
	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/provider/kernel"
	"github.com/lixiaoqing18/zima/framework/provider/setting"
)

func main() {
	framework.Bind(setting.NewZimaSettingProvider(""))

	if engine, err := web.NewWebEngine(); err == nil {
		framework.Bind(kernel.NewZimaGinProvider(engine))
	} else {
		framework.Bind(kernel.NewZimaGinProvider(nil))
	}

	console.RunCommand()
}
