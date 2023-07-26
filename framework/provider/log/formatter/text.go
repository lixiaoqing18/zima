package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/lixiaoqing18/zima/framework/contract"
)

func TextFormatter(level contract.LogLevel, time time.Time, msg string, fields map[string]any) (string, error) {
	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(LevelToStr(level))
	sb.WriteString("]")
	sb.WriteString("\t")
	sb.WriteString(time.Format("2006-01-02 15:04:05"))
	sb.WriteString("\t")
	sb.WriteString(msg)
	sb.WriteString("\t")
	sb.WriteString(fmt.Sprint(fields))
	return sb.String(), nil
}
