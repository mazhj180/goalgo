package stack

type Stack[T any] interface {
	Push(v T)
	Pop() (T, error)
	Peek() (T, error)
	Size() int
	Clear()
}

type node[T any] struct {
	elem    T
	pointer *node[T]
}
