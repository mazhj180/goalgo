package test

import (
	"fmt"
	"mazhj.com/goalgo/data_structures/heap"
	"reflect"
	"testing"
)

var (
	data     = []int{50, 20, 15, 10, 5, 30, 40, 60, 80, 70}
	expected = []int{5, 10, 15, 20, 30, 40, 50, 60, 70, 80}
)

func TestSequenceHeap_InsertAndPeek(t *testing.T) {

	maxHeap := heap.NewSequenceHeap[int](heap.MAX)

	for _, v := range data {
		maxHeap.Insert(v)
	}

	top, err := maxHeap.Peek()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	fmt.Println(top)

}

func TestSequenceHeap_Pop(t *testing.T) {
	minHeap := heap.NewSequenceHeap[int](heap.MIN)

	for _, v := range data {
		minHeap.Insert(v)
	}

	var sorted []int
	for !minHeap.IsEmpty() {
		top, err := minHeap.Pop()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		sorted = append(sorted, top)
	}

	fmt.Println(reflect.DeepEqual(sorted, expected))

}

func TestLinkedHeap_InsertAndPeek(t *testing.T) {

	maxHeap := heap.NewLinkedHeap[int](heap.MAX)

	for _, v := range data {
		maxHeap.Insert(v)
	}

	top, err := maxHeap.Peek()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	fmt.Println(top)

}

func TestLinkedHeap_Pop(t *testing.T) {
	maxHeap := heap.NewLinkedHeap[int](heap.MAX)

	for _, v := range data {
		maxHeap.Insert(v)
	}

	var sorted []int
	for !maxHeap.IsEmpty() {
		top, _ := maxHeap.Pop()
		sorted = append([]int{top}, sorted...)
	}

	fmt.Println(reflect.DeepEqual(sorted, expected))
}
