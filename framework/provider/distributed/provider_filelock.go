package distributed

import (
	"fmt"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaDistributedFileLockProvider struct {
}

func NewZimaDistributedFileLockProviderr() *ZimaDistributedFileLockProvider {
	return &ZimaDistributedFileLockProvider{}
}

func (p *ZimaDistributedFileLockProvider) Name() string {
	return contract.DistributedKey
}
func (p *ZimaDistributedFileLockProvider) FactoryMethod(c framework.Container) framework.NewInstance {
	return NewZimaDistributedFileLockService
}

func (p *ZimaDistributedFileLockProvider) Params(c framework.Container) []any {
	params := []any{c}
	return params
}

func (p *ZimaDistributedFileLockProvider) Lazy() bool {
	return true
}

func (p *ZimaDistributedFileLockProvider) Initialize(c framework.Container) error {
	fmt.Println("ZimaDistributeFileLockProvider begin initializing")
	return nil
}
