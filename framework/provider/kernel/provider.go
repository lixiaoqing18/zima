package kernel

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaGinProvider struct {
	engine *gin.Engine
}

func NewZimaGinProvider(e *gin.Engine) *ZimaGinProvider {
	return &ZimaGinProvider{
		engine: e,
	}
}

func (p *ZimaGinProvider) Name() string {
	return contract.KernelKey
}
func (p *ZimaGinProvider) FactoryMethod(c framework.Container) framework.NewInstance {
	return NewZimaGinService
}

func (p *ZimaGinProvider) Params(c framework.Container) []any {
	params := []any{c, p.engine}
	return params
}

func (p *ZimaGinProvider) Lazy() bool {
	return true
}

func (p *ZimaGinProvider) Initialize(c framework.Container) error {
	fmt.Println("ZimaGinProvider begin initializing")
	if p.engine == nil {
		p.engine = gin.Default()
	}
	return nil
}
