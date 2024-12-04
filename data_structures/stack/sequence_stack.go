package stack

import "errors"

type SequenceStack[T any] []T

func (s *SequenceStack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *SequenceStack[T]) Pop() (T, error) {
	if s.Size() == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}
	st := *s
	top := st[len(st)-1]
	*s = st[:len(st)-1]
	return top, nil
}

func (s SequenceStack[T]) Peek() (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}
	return s[len(s)-1], nil
}

func (s SequenceStack[T]) Size() int {
	return len(s)
}

func (s *SequenceStack[T]) Clear() {
	*s = (*s)[:0]
}
