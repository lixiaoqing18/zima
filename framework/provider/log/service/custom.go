package service

import (
	"io"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaLogCustomService struct {
	ZimaLogService
}

func NewZimaLogCustomService(params ...any) (any, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	formatter := params[2].(contract.Formatter)
	ctxFielder := params[3].(contract.ContextFielder)
	out := params[4].(io.Writer)
	service := &ZimaLogConsoleService{}
	service.container = c
	service.SetContextFielder(ctxFielder)
	service.SetFormatter(formatter)
	service.SetLevel(level)
	service.SetOutput(out)
	return service, nil
}
