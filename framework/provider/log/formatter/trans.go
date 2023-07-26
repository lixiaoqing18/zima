package formatter

import "github.com/lixiaoqing18/zima/framework/contract"

func LevelToStr(level contract.LogLevel) string {
	switch level {
	case contract.PanicLevel:
		return "Panic"
	case contract.FatalLevel:
		return "Fatal"
	case contract.ErrorLevel:
		return "Error"
	case contract.WarnLevel:
		return "Warn"
	case contract.InfoLevel:
		return "Info"
	case contract.DebugLevel:
		return "Debug"
	case contract.TraceLevel:
		return "Trace"
	default:
		return "Info"
	}
}
