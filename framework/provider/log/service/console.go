package service

import (
	"os"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaLogConsoleService struct {
	ZimaLogService
}

func NewZimaLogConsoleService(params ...any) (any, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	formatter := params[2].(contract.Formatter)
	ctxFielder := params[3].(contract.ContextFielder)
	service := &ZimaLogConsoleService{}
	service.container = c
	service.SetContextFielder(ctxFielder)
	service.SetFormatter(formatter)
	service.SetLevel(level)
	service.SetOutput(os.Stdout)
	return service, nil
}
