package list

import "errors"

type doublyListNode[T any] struct {
	elem T
	prev *doublyListNode[T]
	next *doublyListNode[T]
}

type DoublyList[T any] struct {
	head *doublyListNode[T]
	tail *doublyListNode[T]
	size int
}

func NewDoublyList[T any]() *DoublyList[T] {
	return &DoublyList[T]{}
}

func (l *DoublyList[T]) InsertAtHead(elem T) {
	newNode := &doublyListNode[T]{elem: elem}
	if l.head != nil {
		l.head.prev = newNode
	} else {
		l.tail = newNode
	}
	newNode.next = l.head
	l.head = newNode
	l.size++
}

func (l *DoublyList[T]) Insert(elem T) {
	newNode := &doublyListNode[T]{elem: elem}
	if l.tail != nil {
		l.tail.next = newNode
	} else {
		l.head = newNode
	}
	newNode.prev = l.tail
	l.tail = newNode
	l.size++
}

func (l *DoublyList[T]) RemoveHead() error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.size--
	return nil
}

func (l *DoublyList[T]) Remove() error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	l.size--
	return nil
}

func (l *DoublyList[T]) Size() int {
	return l.size
}

func (l *DoublyList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *DoublyList[T]) Traverse(callback Callback[T]) {
	for current := l.head; current != nil; current = current.next {
		callback(current.elem)
	}
}

func (l *DoublyList[T]) TraverseReverse(callback Callback[T]) {
	for current := l.tail; current != nil; current = current.prev {
		callback(current.elem)
	}
}

func (l *DoublyList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}
