package test

import (
	"fmt"
	stack2 "mazhj.com/goalgo/data_structures/stack"
	"testing"
)

func TestSStack(t *testing.T) {
	ss := stack2.SequenceStack[int]{}

	// 测试 Push
	ss.Push(10)
	ss.Push(20)
	ss.Push(30)

	fmt.Println("栈大小:", ss.Size()) // 输出: 栈大小: 3

	// 测试 Peek
	top, _ := ss.Peek()
	fmt.Println("栈顶元素:", top) // 输出: 栈顶元素: 30

	// 测试 Pop
	for ss.Size() > 0 {
		value, _ := ss.Pop()
		fmt.Println("弹出元素:", value)
	}

	// 测试空栈操作
	_, err := ss.Pop()
	if err != nil {
		fmt.Println("错误:", err) // 输出: 错误: ss is empty
	}

	// 测试 Clear
	ss.Push(100)
	ss.Push(200)
	ss.Clear()
	fmt.Println("栈是否为空:", ss.Size()) // 输出: 栈是否为空: 0
}

func TestLStack(t *testing.T) {
	ls := stack2.LinkedStack[int]{}

	// 测试 Push
	ls.Push(10)
	ls.Push(20)
	ls.Push(30)

	fmt.Println("栈大小:", ls.Size()) // 输出: 栈大小: 3

	// 测试 Peek
	top, _ := ls.Peek()
	fmt.Println("栈顶元素:", top) // 输出: 栈顶元素: 30

	// 测试 Pop
	for ls.Size() > 0 {
		value, _ := ls.Pop()
		fmt.Println("弹出元素:", value)
	}

	// 测试空栈操作
	_, err := ls.Pop()
	if err != nil {
		fmt.Println("错误:", err) // 输出: 错误: ls is empty
	}

	// 测试 Clear
	ls.Push(100)
	ls.Push(200)
	ls.Clear()
	fmt.Println("栈是否为空:", ls.Size()) // 输出: 栈是否为空: 0
}
