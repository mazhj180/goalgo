package test

import (
	"fmt"
	"goalgo/data_structures/list"
	"testing"
)

var (
	data     = []int{80, 70, 60, 50, 40, 30, 20, 15, 10, 5}
	expected = []int{5, 10, 15, 20, 30, 40, 50, 60, 70, 80}
)

func TestSinglyList_Head(t *testing.T) {
	sl := list.NewSinglyList[int]()

	// 测试插入到头部
	for _, v := range data {
		sl.InsertAtHead(v)
	}
	sl.Traverse(func(v int) {
		fmt.Printf("%d,", v)
	})
	if sl.Size() != len(data) {
		t.Errorf("Expected size to be 3, got %d", sl.Size())
	}

	// 移除头部
	fmt.Printf("\n移除头部后:")
	err := sl.RemoveHead()
	if err != nil {
		t.Errorf("Error while removing head: %s", err)
	}

	sl.Traverse(func(v int) {
		fmt.Printf("%d,", v)
	})
}

func TestSinglyList_Tail(t *testing.T) {
	sl := list.NewSinglyList[int]()

	for _, v := range data {
		sl.Insert(v)
	}
	sl.Traverse(func(v int) {
		fmt.Printf("%d,", v)
	})

	fmt.Printf("\n移除尾部: ")
	if err := sl.Remove(); err != nil {
		t.Errorf("Error while removing tail: %s", err)
	}
	sl.Traverse(func(v int) {
		fmt.Printf("%d,", v)
	})
}
