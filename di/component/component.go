package component

type Component[T any] struct {
	t *T
}

func Register[T any](_ func() *T) *Component[T] {
	return &Component[T]{}
}

func RegisterE[T any](_ func() *T) *Component[T] {
	return &Component[T]{}
}

func Provider[T any](_ func() *T) *Component[T] {
	return &Component[T]{}
}

func ProviderE[T any](_ func() *T) *Component[T] {
	return &Component[T]{}
}

func (c *Component[T]) Name(name string) *Component[T] {
	return c
}
