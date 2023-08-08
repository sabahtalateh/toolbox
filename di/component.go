package di

type С[T any] struct{}

func Component[T any](_ any) *С[T] {
	return nil
}

func NamedComponent[T any](_ string, _ any) *С[T] {
	return nil
}

func Auto[T any]() T {
	var t T
	return t
}

type Type[T any] struct {
	t T
}

func OfType[T any]() *Type[T] {
	return nil
}

func (t *Type[T]) Named(string, ...string) *Type[T] {
	return nil
}

func (t *Type[T]) First() T {
	return t.t
}

func (t *Type[T]) All() []T {
	return nil
}

func (t *Type[T]) Implementing() []T {
	return nil
}
