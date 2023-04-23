package constructor

type Component[T any] struct {
	t *T
}

func Register[T any](_ any) *Component[T] {
	return &Component[T]{}
}

func (c *Component[T]) Name(name string) *Component[T] {
	return c
}
