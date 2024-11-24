package list

import "errors"

type singlyNode[T any] struct {
	elem T
	next *singlyNode[T]
}

type SinglyList[T any] struct {
	head *singlyNode[T]
	tail *singlyNode[T]
	size int
}

func NewSinglyList[T any]() *SinglyList[T] {
	return &SinglyList[T]{}
}

func (l *SinglyList[T]) InsertAtHead(v T) {
	newNode := &singlyNode[T]{
		elem: v,
		next: l.head,
	}
	if l.tail == nil {
		l.tail = newNode
	}
	l.head = newNode
	l.size++
}

func (l *SinglyList[T]) InsertAtTail(v T) {
	newNode := &singlyNode[T]{
		elem: v,
	}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
	l.size++
}

func (l *SinglyList[T]) RemoveHead() error {
	if l.head == nil {
		return errors.New("list is empty")
	}
	if l.head == l.tail {
		l.tail = nil
	}
	l.head = l.head.next
	l.size--
	return nil
}

func (l *SinglyList[T]) RemoveTail() error {
	if l.tail == nil {
		return errors.New("list is empty")
	}
	if l.tail == l.head {
		l.head = nil
		l.tail = nil
	} else {
		tailBefore := l.head
		for tailBefore.next != l.tail {
			tailBefore = tailBefore.next
		}
		tailBefore.next = nil
		l.tail = tailBefore
	}
	l.size--
	return nil
}

func (l *SinglyList[T]) Traverse(callback Callback[T]) {
	for current := l.head; current != nil; current = current.next {
		callback(current.elem)
	}
}

func (l *SinglyList[T]) Size() int {
	return l.size
}

func (l *SinglyList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *SinglyList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}
