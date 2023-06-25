package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lixiaoqing18/zima/app/web/module/demo"
)

func RegisterRouter(core *gin.Engine) {
	core.Static("/dist/", "./dist/")

	demo.Register(core)
}
