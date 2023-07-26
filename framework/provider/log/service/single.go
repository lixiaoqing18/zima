package service

import (
	"os"
	"path/filepath"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaLogSingleService struct {
	ZimaLogService
}

func NewZimaLogSingleService(params ...any) (any, error) {
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
	folder := configService.GetString("log.single.folder")
	if folder == "" {
		folder = settingService.LogFolder()
	}
	filename := "zima.log"
	if configService.IsExist("log.single.filename") {
		filename = configService.GetString("log.single.filename")
	}
	f, err := os.OpenFile(filepath.Join(folder, filename), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	service.SetOutput(f)
	return service, nil
}
