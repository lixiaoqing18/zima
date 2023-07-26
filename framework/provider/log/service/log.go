package service

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaLogService struct {
	container  framework.Container
	level      contract.LogLevel
	formatter  contract.Formatter
	ctxFielder contract.ContextFielder
	out        io.Writer
}

func (service *ZimaLogService) Panic(c context.Context, msg string, fields map[string]any) {
	service.logF(contract.PanicLevel, c, msg, fields)
}

func (service *ZimaLogService) Fatal(c context.Context, msg string, fields map[string]any) {
	service.logF(contract.FatalLevel, c, msg, fields)
}

func (service *ZimaLogService) Error(c context.Context, msg string, fields map[string]any) {
	service.logF(contract.ErrorLevel, c, msg, fields)
}
func (service *ZimaLogService) Warn(c context.Context, msg string, fields map[string]any) {
	service.logF(contract.WarnLevel, c, msg, fields)
}

func (service *ZimaLogService) Info(c context.Context, msg string, fields map[string]any) {
	service.logF(contract.InfoLevel, c, msg, fields)
}

func (service *ZimaLogService) Debug(c context.Context, msg string, fields map[string]any) {
	service.logF(contract.DebugLevel, c, msg, fields)
}

func (service *ZimaLogService) Trace(c context.Context, msg string, fields map[string]any) {
	service.logF(contract.TraceLevel, c, msg, fields)
}

func (service *ZimaLogService) SetLevel(level contract.LogLevel) {
	service.level = level
}

func (service *ZimaLogService) SetContextFielder(f contract.ContextFielder) {
	service.ctxFielder = f
}

func (service *ZimaLogService) SetFormatter(f contract.Formatter) {
	service.formatter = f
}

func (service *ZimaLogService) SetOutput(writer io.Writer) {
	service.out = writer
}

func (service *ZimaLogService) logF(level contract.LogLevel, c context.Context, msg string, fields map[string]any) {
	if service.level < level {
		return
	}
	if fields == nil {
		fields = map[string]any{}
	}
	if service.ctxFielder != nil {
		ctxFields := service.ctxFielder(c)
		for k, v := range ctxFields {
			fields[k] = v
		}
	}
	//todo TraceLevel的处理

	contents, err := service.formatter(level, time.Now(), msg, fields)
	if err != nil {
		return
	}

	io.WriteString(service.out, contents)
	io.WriteString(service.out, "\n")

	if level == contract.PanicLevel {
		panic(contents)
	}

	if level == contract.FatalLevel {
		//todo fatal 时的钩子FatalHandler，用于优雅关停
		os.Exit(1)
	}
}
