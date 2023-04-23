package request

type Header[T any] struct {
	v *T
}

func (h *Header[T]) Get() *T {
	return h.v
}

type PathParam[T any] struct {
	v *T
}

func (h *PathParam[T]) Get() *T {
	return h.v
}

type QueryParam[T any] struct {
	v *T
}

func (h *QueryParam[T]) Get() *T {
	return h.v
}

type Body[T any] struct {
	v *T
}

func (h *Body[T]) Get() *T {
	return h.v
}
