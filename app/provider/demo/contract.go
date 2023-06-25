package demo

const Key = "zima:demo"

type IService interface {
	Demo() Foo
}

type Foo struct {
	Message string
	Code    int
}
