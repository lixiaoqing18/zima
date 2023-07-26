package log

import (
	"fmt"
	"io"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
	"github.com/lixiaoqing18/zima/framework/provider/log/formatter"
	"github.com/lixiaoqing18/zima/framework/provider/log/service"
)

type ZimaLogProvider struct {
	Driver     string
	Level      contract.LogLevel
	Out        io.Writer
	Formatter  contract.Formatter
	CtxFielder contract.ContextFielder
}

func NewZimaLogProvider() *ZimaLogProvider {
	return &ZimaLogProvider{}
}

func (p *ZimaLogProvider) Name() string {
	return contract.LogKey
}

func (p *ZimaLogProvider) FactoryMethod(c framework.Container) framework.NewInstance {
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	driver := ""
	if p.Driver != "" {
		driver = p.Driver
	} else {
		driver = configService.GetString("log.driver")
	}

	if driver == "custom" {
		return service.NewZimaLogCustomService
	}

	levelStr := configService.GetString("log.level")
	p.Level = p.toLevel(levelStr)

	formatterStr := configService.GetString("log.formatter")
	switch formatterStr {
	case "text":
		p.Formatter = formatter.TextFormatter
	case "json":
		p.Formatter = formatter.JsonFormatter
	default:
		p.Formatter = formatter.TextFormatter
	}

	switch driver {
	case "console":
		return service.NewZimaLogConsoleService
	case "single":
		return service.NewZimaLogSingleService
	case "rotate":
		return service.NewZimaLogRotateService
	case "custom":
		return service.NewZimaLogCustomService
	default:
		return service.NewZimaLogConsoleService
	}
}

func toLevel(levelStr string) {
	panic("unimplemented")
}

func (p *ZimaLogProvider) Params(c framework.Container) []any {
	params := []any{c, p.Level, p.Formatter, p.CtxFielder, p.Out}
	return params
}

func (p *ZimaLogProvider) Lazy() bool {
	return true
}

func (p *ZimaLogProvider) Initialize(c framework.Container) error {
	fmt.Println("ZimaLogProvider begin initializing")
	return nil
}

func (p *ZimaLogProvider) toLevel(level string) contract.LogLevel {
	switch level {
	case "panic":
		return contract.PanicLevel
	case "fatal":
		return contract.FatalLevel
	case "error":
		return contract.ErrorLevel
	case "warn":
		return contract.WarnLevel
	case "info":
		return contract.InfoLevel
	case "debug":
		return contract.DebugLevel
	case "trace":
		return contract.TraceLevel
	default:
		return contract.InfoLevel
	}
}
