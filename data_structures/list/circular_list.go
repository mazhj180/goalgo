package list

import (
	"errors"
)

type circularListNode[T any] struct {
	elem T
	next *circularListNode[T]
}

type CircularList[T any] struct {
	head *circularListNode[T]
	size int
}

func NewCircularList[T any]() *CircularList[T] {
	return &CircularList[T]{}
}

func (l *CircularList[T]) Insert(elem T) {
	newNode := &circularListNode[T]{elem: elem}
	if l.head == nil {
		l.head = newNode
		newNode.next = newNode // 自己指向自己形成环
	} else {
		current := l.head
		for current.next != l.head { // 找到最后一个节点
			current = current.next
		}
		current.next = newNode
		newNode.next = l.head // 新节点指向头部，形成环
	}
	l.size++
}

func (l *CircularList[T]) Remove() error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	if l.size == 1 { // 只有一个节点的情况
		l.head = nil
		l.size = 0
		return nil
	}
	current := l.head
	for current.next != l.head {
		current = current.next
	}
	current.next = l.head.next
	l.head = l.head.next
	l.size--
	return nil
}

func (l *CircularList[T]) Traverse(callback Callback[T]) {
	current := l.head
	for {
		callback(current.elem)
		current = current.next
		if current == l.head {
			break
		}
	}
}

func (l *CircularList[T]) Size() int {
	return l.size
}

func (l *CircularList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *CircularList[T]) Clear() {
	if l.IsEmpty() {
		return
	}

	current := l.head
	for {
		next := current.next
		current.next = nil // 断开引用
		if next == l.head {
			break
		}
		current = next
	}
	l.head = nil
	l.size = 0
}
