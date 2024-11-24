package heap

import (
	"cmp"
	"errors"
)

type LinkedHeap[T cmp.Ordered] struct {
	root *node[T]
	tail *node[T]
	size int
	attr Attr
}

func NewLinkedHeap[T cmp.Ordered](attr Attr) *LinkedHeap[T] {
	return &LinkedHeap[T]{
		attr: attr,
	}
}

func (h *LinkedHeap[T]) upHeap(n *node[T]) {

	for n.parent != nil {

		if h.attr == MAX && n.elem <= n.parent.elem {
			break
		}

		if h.attr == MIN && n.elem >= n.parent.elem {
			break
		}

		n.elem, n.parent.elem = n.parent.elem, n.elem

		n = n.parent
	}

}

func (h *LinkedHeap[T]) downHeap(n *node[T]) {

	for n != nil {

		best := n
		if h.attr == MAX {
			if n.left != nil && n.left.elem > best.elem {
				best = n.left
			}

			if n.right != nil && n.right.elem > best.elem {
				best = n.right
			}
		}

		if h.attr == MIN {
			if n.left != nil && n.left.elem < best.elem {
				best = n.left
			}

			if n.right != nil && n.right.elem < best.elem {
				best = n.right
			}
		}

		if best == n {
			break
		}

		best.elem, n.elem = n.elem, best.elem
		n = best
	}
}

func (h *LinkedHeap[T]) findParent(pos int) *node[T] {
	path := h.binaryPath(pos >> 1)
	return h.navigate(path)
}

func (h *LinkedHeap[T]) findTail(pos int) *node[T] {
	if pos == 0 {
		return nil
	}
	path := h.binaryPath(pos)
	return h.navigate(path)
}

func (h *LinkedHeap[T]) binaryPath(pos int) []bool {
	var path []bool

	for pos > 1 {
		path = append([]bool{pos%2 == 1}, path...)
		pos /= 2
	}
	return path
}

func (h *LinkedHeap[T]) navigate(path []bool) *node[T] {
	cur := h.root

	for _, p := range path {
		if !p {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}

	return cur
}

func (h *LinkedHeap[T]) removeTail() {
	if h.tail == nil {
		return
	}
	parent := h.tail.parent
	if parent != nil {
		if parent.left == h.tail {
			parent.left = nil
		} else if parent.right == h.tail {
			parent.right = nil
		}
	}
	h.tail.parent = nil
	h.tail = h.findTail(h.size - 1)
}

func (h *LinkedHeap[T]) Insert(v T) {
	newNode := &node[T]{
		elem: v,
	}

	if h.root == nil {
		h.root = newNode
		h.tail = newNode
		h.size = 1
		return
	}

	parent := h.findParent(h.size + 1)
	if parent.left == nil {
		parent.left = newNode
	} else if parent.right == nil {
		parent.right = newNode
	}

	newNode.parent = parent
	h.tail = newNode
	h.size += 1
	h.upHeap(newNode)

}

func (h *LinkedHeap[T]) Pop() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, errors.New("empty heap")
	}

	top := h.root.elem
	h.root.elem = h.tail.elem
	h.removeTail()
	h.size -= 1
	h.downHeap(h.root)

	return top, nil
}

func (h *LinkedHeap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, errors.New("empty heap")
	}
	return h.root.elem, nil
}

func (h *LinkedHeap[T]) IsEmpty() bool {
	return h.size == 0
}

func (h *LinkedHeap[T]) Size() int {
	return h.size
}
