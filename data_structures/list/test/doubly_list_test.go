package test

import (
	"fmt"
	"mazhj.com/goalgo/data_structures/list"
	"testing"
)

func TestDoublyListInsert(t *testing.T) {

	dl := list.NewDoublyList[int]()

	// 头插入
	for _, v := range data {
		dl.InsertAtHead(v)
	}

	// 顺序遍历
	fmt.Print("顺序遍历:")
	dl.Traverse(func(v int) {
		fmt.Printf("%d,", v)
	})

	l := dl.Size()
	for i := 0; i < l; i++ {
		_ = dl.RemoveHead()
	}

	fmt.Printf("\n")

	for _, v := range data {
		dl.InsertAtTail(v)
	}
	// 逆序遍历
	fmt.Print("逆序遍历 : ")
	dl.TraverseReverse(func(v int) {
		fmt.Printf("%d,", v)
	})
	fmt.Println()

	_ = dl.RemoveTail()
	fmt.Print("删除尾元素后：")
	dl.Traverse(func(v int) {
		fmt.Printf("%d,", v)
	})

}
