package stack

import "errors"

type SequenceStack[T any] struct {
	elems []T
}

func (s *SequenceStack[T]) Push(v T) {
	s.elems = append(s.elems, v)
}

func (s *SequenceStack[T]) Pop() (T, error) {
	if len(s.elems) == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem, nil
}

func (s *SequenceStack[T]) Peek() (T, error) {
	if len(s.elems) == 0 {
		var zero T
		return zero, errors.New("empty stack")
	}
	return s.elems[len(s.elems)-1], nil
}

func (s *SequenceStack[T]) Size() int {
	return len(s.elems)
}

func (s *SequenceStack[T]) Clear() {
	s.elems = nil
}
