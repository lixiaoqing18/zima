package setting

import (
	"errors"
	"path/filepath"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/util"

	"github.com/google/uuid"
)

type ZimaSettingService struct {
	container  framework.Container
	baseFolder string
	appID      string
}

func NewZimaSettingService(params ...any) (any, error) {
	if len(params) != 2 {
		return nil, errors.New("params length must be 2")
	}
	c := params[0].(framework.Container)
	bf := params[1].(string)
	id := uuid.New().String()
	return &ZimaSettingService{
		container:  c,
		baseFolder: bf,
		appID:      id,
	}, nil
}

func (app *ZimaSettingService) AppID() string {
	return app.appID
}

// Version 定义当前版本
func (app *ZimaSettingService) Version() string {
	return "0.1"
}

// BaseFolder 定义项目基础地址
func (app *ZimaSettingService) BaseFolder() string {
	//优先取参数传递的baseFolder
	if app.baseFolder != "" {
		return app.baseFolder
	}
	//其次取命令行参数的base_folder
	/*
		var baseFolder string
		if flag.Lookup("base_folder") == nil {
			flag.StringVar(&baseFolder, "base_folder", "", "base_folder 参数, 默认为当前路径")
			flag.Parse()
		}
		if baseFolder != "" {
			return baseFolder
		}
	*/
	//最后取当前程序运行的目录
	return util.GetExecDirectory()
}

// ConfigFolder 定义了配置文件的路径
func (app *ZimaSettingService) ConfigFolder() string {
	return filepath.Join(app.BaseFolder(), "config")
}

// LogFolder 定义了日志所在路径
func (app *ZimaSettingService) LogFolder() string {
	return filepath.Join(app.BaseFolder(), "storage", "log")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (app *ZimaSettingService) ProviderFolder() string {
	return filepath.Join(app.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (app *ZimaSettingService) MiddlewareFolder() string {
	return filepath.Join(app.BaseFolder(), "web", "middleware")
}

// CommandFolder 定义业务定义的命令
func (app *ZimaSettingService) CommandFolder() string {
	return filepath.Join(app.BaseFolder(), "console", "cmd")
}

// RuntimeFolder 定义业务的运行中间态信息
func (app *ZimaSettingService) RuntimeFolder() string {
	return filepath.Join(app.BaseFolder(), "storage", "runtime")
}

// TestFolder 存放测试所需要的信息
func (app *ZimaSettingService) TestFolder() string {
	return filepath.Join(app.BaseFolder(), "test")
}
