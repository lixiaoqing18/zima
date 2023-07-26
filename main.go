package main

import (
	"context"

	"github.com/lixiaoqing18/zima/app/console"
	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/provider/config"
	"github.com/lixiaoqing18/zima/framework/provider/distributed"
	"github.com/lixiaoqing18/zima/framework/provider/env"
	"github.com/lixiaoqing18/zima/framework/provider/log"
	"github.com/lixiaoqing18/zima/framework/provider/setting"
)

func main() {
	framework.Bind(setting.NewZimaSettingProvider(""))
	framework.Bind(env.NewZimaEnvProvider())
	framework.Bind(config.NewZimaConfigProvider())
	framework.Bind(distributed.NewZimaDistributedFileLockProviderr())
	framework.Bind(&log.ZimaLogProvider{
		CtxFielder: func(ctx context.Context) map[string]any {
			result := map[string]any{}
			result["traceId"] = ctx.Value("traceId")
			return result
		},
		/*
			Level:     contract.InfoLevel,
			Formatter: formatter.JsonFormatter,
			Driver:    "custom",
			Out:       os.Stdout,
		*/
	})
	/*
		if engine, err := web.NewWebEngine(); err == nil {
			framework.Bind(kernel.NewZimaGinProvider(engine))
		} else {
			framework.Bind(kernel.NewZimaGinProvider(nil))
		}
	*/

	console.RunCommand()
}
