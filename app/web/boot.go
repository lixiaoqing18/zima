package web

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
	"github.com/lixiaoqing18/zima/framework/middleware"
)

func NewWebEngine() (*gin.Engine, error) {
	app := framework.MustMake(contract.SettingKey).(contract.Setting)
	f, _ := os.Create(filepath.Join(app.LogFolder(), "gin.log"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	core := gin.New()
	core.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	//core.Use(gin.Logger())
	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	RegisterRouter(core)

	return core, nil
}
