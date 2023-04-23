package response

type Response[T any] struct {
	X *T
}

type Ok[T any] struct {
	X *T
}

func NewOk[T any](x *T) *Ok[T] {
	return &Ok[T]{X: x}
}

type NotFound[T any] struct {
	X *T
}

func NewNotFound[T any](x *T) *NotFound[T] {
	return &NotFound[T]{X: x}
}

type InternalServerError[T any] struct {
	X *T
}

func NewInternalServerError[T any](x *T) *InternalServerError[T] {
	return &InternalServerError[T]{X: x}
}
