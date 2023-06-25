package demo

import (
	"github.com/lixiaoqing18/zima/framework"
)

type Service struct {
	IService
	c framework.Container
}

func NewDemoService(params ...any) (any, error) {
	container := params[0].(framework.Container)
	service := &Service{
		c: container,
	}
	return service, nil
}

func (service *Service) Demo() Foo {
	return Foo{
		Message: "this is a foo",
		Code:    200,
	}
}
