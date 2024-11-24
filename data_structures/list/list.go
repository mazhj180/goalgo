package list

type Callback[T any] func(v T)

type LinkedList[T any] interface {
	IsEmpty() bool
	Size() int
	InsertAtHead(T)
	InsertAtTail(T)
	RemoveHead() error
	RemoveTail() error
	Traverse(Callback[T])
	Clear()
}
