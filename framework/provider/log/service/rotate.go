package service

import (
	"path/filepath"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZimaLogRotateService struct {
	ZimaLogService
}

func NewZimaLogRotateService(params ...any) (any, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	formatter := params[2].(contract.Formatter)
	ctxFielder := params[3].(contract.ContextFielder)
	service := &ZimaLogConsoleService{}
	service.container = c
	service.SetContextFielder(ctxFielder)
	service.SetFormatter(formatter)
	service.SetLevel(level)

	settingService := framework.MustMake(contract.SettingKey).(contract.Setting)
	configService := framework.MustMake(contract.ConfigKey).(contract.Config)
	folder := configService.GetString("log.rotate.folder")
	if folder == "" {
		folder = settingService.LogFolder()
	}
	filename := "zima_rotate.log"
	if configService.IsExist("log.rotate.filename") {
		filename = configService.GetString("log.rotate.filename")
	}
	maxsize := 500
	if configService.IsExist("log.rotate.maxsize") {
		maxsize = configService.GetInt("log.rotate.maxsize")
	}
	maxbackups := 3
	if configService.IsExist("log.rotate.maxbackups") {
		maxbackups = configService.GetInt("log.rotate.maxbackups")
	}
	maxage := 7
	if configService.IsExist("log.rotate.maxage") {
		maxbackups = configService.GetInt("log.rotate.maxage")
	}
	compress := false
	if configService.IsExist("log.rotate.compress") {
		compress = configService.GetBool("log.rotate.compress")
	}
	logWriter := &lumberjack.Logger{
		Filename:   filepath.Join(folder, filename),
		MaxSize:    maxsize, // megabytes
		MaxBackups: maxbackups,
		MaxAge:     maxage,   //days
		Compress:   compress, // disabled by default
	}
	service.SetOutput(logWriter)
	return service, nil
}
