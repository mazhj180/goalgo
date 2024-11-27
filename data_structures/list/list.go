package list

type Callback[T any] func(v T)

type LinkedList[T any] interface {
	IsEmpty() bool
	Size() int
	Insert(T)
	Remove() error
	Traverse(Callback[T])
	Clear()
}
