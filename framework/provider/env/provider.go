package env

import (
	"fmt"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaEnvProvider struct {
}

func NewZimaEnvProvider() *ZimaEnvProvider {
	return &ZimaEnvProvider{}
}

func (p *ZimaEnvProvider) Name() string {
	return contract.EnvKey
}
func (p *ZimaEnvProvider) FactoryMethod(c framework.Container) framework.NewInstance {
	return NewZimaEnvService
}

func (p *ZimaEnvProvider) Params(c framework.Container) []any {
	settingService := c.MustMake(contract.SettingKey).(contract.Setting)
	params := []any{c, settingService.BaseFolder()}
	return params
}

func (p *ZimaEnvProvider) Lazy() bool {
	return false
}

func (p *ZimaEnvProvider) Initialize(c framework.Container) error {
	fmt.Println("ZimaEnvProvider begin initializing")
	return nil
}
