package stack

import "errors"

type LinkedStack[T any] struct {
	top  *node[T]
	size int
}

func (l *LinkedStack[T]) Push(v T) {
	newNode := &node[T]{elem: v, pointer: l.top}
	l.top = newNode
	l.size++
}

func (l *LinkedStack[T]) Pop() (T, error) {
	if l.top == nil {
		var zero node[T]
		return zero.elem, errors.New("empty stack")
	}
	elem := l.top.elem
	l.top = l.top.pointer
	l.size--
	return elem, nil
}

func (l *LinkedStack[T]) Peek() (T, error) {
	if l.top == nil {
		var zero node[T]
		return zero.elem, errors.New("empty stack")
	}
	return l.top.elem, nil
}

func (l *LinkedStack[T]) Size() int {
	return l.size
}

func (l *LinkedStack[T]) Clear() {
	l.top = nil
	l.size = 0
}
