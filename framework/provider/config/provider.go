package config

import (
	"fmt"
	"path/filepath"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaConfigProvider struct {
}

func NewZimaConfigProvider() *ZimaConfigProvider {
	return &ZimaConfigProvider{}
}

func (p *ZimaConfigProvider) Name() string {
	return contract.ConfigKey
}
func (p *ZimaConfigProvider) FactoryMethod(c framework.Container) framework.NewInstance {
	return NewZimaConfigService
}

func (p *ZimaConfigProvider) Params(c framework.Container) []any {
	settingService := c.MustMake(contract.SettingKey).(contract.Setting)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	configFolder := filepath.Join(settingService.ConfigFolder(), envService.AppEnv())
	params := []any{c, configFolder, envService.All()}
	return params
}

func (p *ZimaConfigProvider) Lazy() bool {
	return false
}

func (p *ZimaConfigProvider) Initialize(c framework.Container) error {
	fmt.Println("ZimaConfigProvider begin initializing")
	return nil
}
