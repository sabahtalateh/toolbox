package di

func Storage[T any]() {}

func FromStorage[T any]() T {
	var t T
	return t
}
