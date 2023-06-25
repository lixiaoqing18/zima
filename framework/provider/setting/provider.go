package setting

import (
	"fmt"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaSettingProvider struct {
	BaseFolder string
}

func NewZimaSettingProvider(baseFolder string) *ZimaSettingProvider {
	return &ZimaSettingProvider{
		BaseFolder: baseFolder,
	}
}

func (p *ZimaSettingProvider) Name() string {
	return contract.SettingKey
}
func (p *ZimaSettingProvider) FactoryMethod(c framework.Container) framework.NewInstance {
	return NewZimaSettingService
}

func (p *ZimaSettingProvider) Params(c framework.Container) []any {
	params := []any{c, p.BaseFolder}
	return params
}

func (p *ZimaSettingProvider) Lazy() bool {
	return true
}

func (p *ZimaSettingProvider) Initialize(c framework.Container) error {
	fmt.Println("ZimaAppProvider beigin initializing")
	return nil
}
