package di

type Container interface {
	Resolve(name string) interface{}
	Clean() error
}
