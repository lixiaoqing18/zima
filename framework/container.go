package framework

import (
	"errors"
	"fmt"
	"sync"

	"github.com/lixiaoqing18/zima/framework/contract"
)

type Container interface {
	Bind(provider ServiceProvider) error
	IsBind(key string) bool
	Make(key string) (any, error)
	MustMake(key string) any
	MakeNew(key string, params []any) (any, error)
}

type ZimaContainer struct {
	Container
	providers map[string]ServiceProvider
	instances map[string]any
	mutex     *sync.RWMutex
}

/*
	func NewZimaContainer() *ZimaContainer {
		return &ZimaContainer{
			providers: map[string]ServiceProvider{},
			instances: map[string]any{},
			mutex:     &sync.RWMutex{},
		}
	}
*/

var zimaContainerInstace Container
var lock = &sync.Mutex{}

func GetContainer() Container {
	if zimaContainerInstace == nil {
		lock.Lock()
		if zimaContainerInstace == nil {
			zimaContainerInstace = &ZimaContainer{
				providers: map[string]ServiceProvider{},
				instances: map[string]any{},
				mutex:     &sync.RWMutex{},
			}
		}
		lock.Unlock()
	}
	return zimaContainerInstace
}

func Bind(provider ServiceProvider) error {
	return GetContainer().Bind(provider)
}

func (c *ZimaContainer) Bind(provider ServiceProvider) error {
	c.mutex.Lock()

	//if c.IsBind(provider.Name()) {
	//	return nil
	//}
	c.providers[provider.Name()] = provider
	c.mutex.Unlock()
	if !provider.Lazy() {
		_, err := c.createInstance(provider, false, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ZimaContainer) IsBind(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, ok := c.providers[key]
	return ok
}

func (c *ZimaContainer) Make(key string) (any, error) {
	return c.make(key, false, nil)
}

func MustMake(key string) any {
	return GetContainer().MustMake(key)
}

func (c *ZimaContainer) MustMake(key string) any {
	ins, err := c.Make(key)
	if err != nil {
		panic(err)
	}
	return ins
}

func (c *ZimaContainer) MakeNew(key string, params []any) (any, error) {
	return c.make(key, true, params)
}

func (c *ZimaContainer) make(key string, new bool, params []any) (any, error) {
	c.mutex.RLock()
	c.mutex.RUnlock()
	if !c.IsBind(key) {
		return nil, errors.New(fmt.Sprintf("provider %s not bind", key))
	}
	if !new {
		if v, ok := c.instances[key]; ok {
			return v, nil
		}
	}

	provider := c.providers[key]

	instance, err := c.createInstance(provider, false, params)

	return instance, err
}

func (c *ZimaContainer) createInstance(provider ServiceProvider, useNewParam bool, newParams []any) (any, error) {
	if err := provider.Initialize(c); err != nil {
		return nil, err
	}
	method := provider.FactoryMethod(c)
	var params []any
	if useNewParam {
		params = newParams
	} else {
		params = provider.Params(c)
	}
	instance, err := method(params...)
	if err != nil {
		return nil, err
	}
	c.instances[provider.Name()] = instance
	return instance, nil
}

var zimaLogInstace contract.Log

func GetLog() contract.Log {
	if zimaLogInstace == nil {
		zimaLogInstace = MustMake(contract.LogKey).(contract.Log)
	}
	return zimaLogInstace
}
