package component

type DiComponent[T any] struct {
	t *T
}

func Component[T any](_ func() *T) *DiComponent[T] {
	return &DiComponent[T]{}
}

func Provider[T any](_ func() *T) *DiComponent[T] {
	return &DiComponent[T]{}
}

func (c *DiComponent[T]) Name(name string) *DiComponent[T] {
	return c
}
