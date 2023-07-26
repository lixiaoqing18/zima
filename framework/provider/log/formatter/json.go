package formatter

import (
	"encoding/json"
	"time"

	"github.com/lixiaoqing18/zima/framework/contract"
)

func JsonFormatter(level contract.LogLevel, time time.Time, msg string, fields map[string]any) (string, error) {
	fields["level"] = LevelToStr(level)
	fields["time"] = time.Format("2006-01-02 15:04:05")
	fields["msg"] = msg
	bytes, err := json.Marshal(fields)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
