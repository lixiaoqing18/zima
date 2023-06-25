package demo

import (
	"fmt"

	"github.com/lixiaoqing18/zima/framework"
)

type ServiceProvider struct {
	framework.ServiceProvider
}

func NewServiceProvider() framework.ServiceProvider {
	return &ServiceProvider{}
}

func (p *ServiceProvider) Name() string {
	return Key
}
func (p *ServiceProvider) FactoryMethod(c framework.Container) framework.NewInstance {
	return NewDemoService
}

func (p *ServiceProvider) Params(c framework.Container) []any {
	params := []any{c}
	return params
}

func (p *ServiceProvider) Lazy() bool {
	return true
}

func (p *ServiceProvider) Initialize(c framework.Container) error {
	fmt.Println("demo provider beigin initializing")
	return nil
}
