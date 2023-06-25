package framework

type NewInstance func(...any) (any, error)

type ServiceProvider interface {
	Name() string
	FactoryMethod(c Container) NewInstance
	Params(c Container) []any
	Lazy() bool
	Initialize(c Container) error
}
