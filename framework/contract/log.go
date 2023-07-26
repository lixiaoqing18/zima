package contract

import (
	"context"
	"io"
	"time"
)

const LogKey = "zima:log"

type LogLevel uint

const (
	FatalLevel LogLevel = iota
	PanicLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type ContextFielder func(ctx context.Context) map[string]any

type Formatter func(level LogLevel, time time.Time, msg string, fields map[string]any) (string, error)

type Log interface {
	Fatal(c context.Context, msg string, fields map[string]any)
	Panic(c context.Context, msg string, fields map[string]any)
	Error(c context.Context, msg string, fields map[string]any)
	Warn(c context.Context, msg string, fields map[string]any)
	Info(c context.Context, msg string, fields map[string]any)
	Debug(c context.Context, msg string, fields map[string]any)
	Trace(c context.Context, msg string, fields map[string]any)

	SetLevel(level LogLevel)
	SetContextFielder(f ContextFielder)
	SetFormatter(f Formatter)
	SetOutput(writer io.Writer)
}
