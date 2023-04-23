package inject

type Config[T any] struct {
	val *T
}

func (i *Config[T]) Get() *T {
	return i.val
}

func (i *Config[T]) ByName(name string) *T {
	return i.val
}

type Struct[T any] struct {
	val *T
}

func (i *Struct[T]) Get() *T {
	return i.val
}

func (i *Struct[T]) ByName(name string) *T {
	return i.val
}

type StructList[T any] struct {
	val []*T
}

func (i *StructList[T]) All() []*T {
	return i.val
}

func (i *StructList[T]) ByName(names ...string) []*T {
	return i.val
}

type Impl[T any] struct {
	val T
}

func (i *Impl[T]) Get() T {
	return i.val
}

func (i *Impl[T]) ByName() T {
	return i.val
}

type ImplList[T any] struct {
	val []T
}

func (i *ImplList[T]) Get() []T {
	return i.val
}

func (i *ImplList[T]) ByName(names ...string) []T {
	return i.val
}
