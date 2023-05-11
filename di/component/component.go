package component

type DiComponent[T any] struct {
	t *T
}

func Component[T any](_ any) *DiComponent[T] {
	return &DiComponent[T]{}
}

func Provider[T any](_ any) *DiComponent[T] {
	return &DiComponent[T]{}
}

func (c *DiComponent[T]) Name(_ string) *DiComponent[T] {
	return c
}
