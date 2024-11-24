package heap

import "cmp"

const (
	MAX Attr = iota
	MIN
)

type Attr int

type Heap[T cmp.Ordered] interface {
	Insert(v T)
	Pop() (T, error)
	Peek() (T, error)
	Size() int
	IsEmpty() bool
}

type node[T cmp.Ordered] struct {
	elem   T
	parent *node[T]
	left   *node[T]
	right  *node[T]
}
