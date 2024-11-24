package heap

import (
	"cmp"
	"errors"
)

type SequenceHeap[T cmp.Ordered] struct {
	elems []T
	size  int
	attr  Attr
}

func NewSequenceHeap[T cmp.Ordered](attr Attr) *SequenceHeap[T] {
	return &SequenceHeap[T]{
		attr: attr,
	}
}

// 逐步构建上浮操作
func (h *SequenceHeap[T]) upHeap(k int) {

	for k > 0 {
		parent := (k - 1) >> 1

		if h.attr == MAX && h.elems[k] <= h.elems[parent] {
			break
		}

		if h.attr == MIN && h.elems[k] >= h.elems[parent] {
			break
		}

		h.elems[k], h.elems[parent] = h.elems[parent], h.elems[k]
		k = parent
	}
}

// 批量插入下沉操作
func (h *SequenceHeap[T]) downHeap(k int) {

	n := h.size

	for k < n>>1 {

		left := 2*k + 1
		right := 2*k + 2

		if h.attr == MAX {
			if right < n && h.elems[left] < h.elems[right] {
				left = right
			}

			if h.elems[k] >= h.elems[left] {
				break
			}
		}

		if h.attr == MIN {
			if right < n && h.elems[left] > h.elems[right] {
				left = right
			}

			if h.elems[k] <= h.elems[left] {
				break
			}
		}

		h.elems[k], h.elems[left] = h.elems[left], h.elems[k]
		k = left
	}

}

func (h *SequenceHeap[T]) Insert(e T) {
	h.elems = append(h.elems, e)
	h.size += 1
	h.upHeap(h.size - 1)
}

func (h *SequenceHeap[T]) Pop() (T, error) {

	if h.IsEmpty() {
		var zero T
		return zero, errors.New("empty heap")
	}
	top := h.elems[0]
	h.elems[0], h.elems[h.size-1] = h.elems[h.size-1], h.elems[0]
	h.elems = h.elems[0 : h.size-1]
	h.size -= 1
	h.downHeap(0)

	return top, nil
}

func (h *SequenceHeap[T]) Peek() (T, error) {

	if h.IsEmpty() {
		var zero T
		return zero, errors.New("empty heap")
	}
	top := h.elems[0]
	return top, nil
}

func (h *SequenceHeap[T]) Size() int {
	return h.size
}

func (h *SequenceHeap[T]) IsEmpty() bool {
	return h.size == 0
}
