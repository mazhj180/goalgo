package test

import (
	cl "container/list"
	"fmt"
	"goalgo/data_structures/list"
	"math/rand"
	"testing"
)

func TestSkipList(t *testing.T) {
	sl := list.NewSkipList()

	for i := 0; i < 100000; i++ {
		score := rand.Float64() * 1000
		sl.Insert(fmt.Sprintf("Item-%d", i), int(score))
	}

	// 打印跳表的结构
	fmt.Println("Skip List after inserting 1000 items:")
	list.Print(sl)

	// 打印层级数
	fmt.Printf("层级数 : %d\n", sl.Level())

}

const testSize = 100000 // 测试数据量

func BenchmarkSkipList(b *testing.B) {
	sl := list.NewSkipList()

	// 插入数据
	insertSkipList(sl)

	// 查找操作
	b.ResetTimer() // 重置计时器，确保只测量查询的时间
	for i := 0; i < b.N; i++ {
		score := rand.Float64() * 1000
		sl.Search(int(score))
	}

	// 删除操作
	//for i := 0; i < b.N; i++ {
	//	score := rand.Int() * 1000
	//	sl.Remove(score)
	//}
}

// 测试 Go 标准库链表的插入、查找和删除
func BenchmarkStdList(b *testing.B) {
	stdList := cl.New()

	// 插入数据
	insertStdList(stdList)

	// 查找操作（遍历）
	b.ResetTimer() // 重置计时器，确保只测量查询的时间
	for i := 0; i < b.N; i++ {
		score := rand.Float64() * 1000
		for e := stdList.Front(); e != nil; e = e.Next() {
			// 模拟查找操作
			if e.Value == int(score) {
				break
			}
		}
	}

	// 删除操作
	//for i := 0; i < b.N; i++ {
	//	score := rand.Float64() * 1000
	//	for e := stdList.Front(); e != nil; e = e.Next() {
	//		if e.Value == fmt.Sprintf("Item-%d", int(score)) {
	//			stdList.Remove(e)
	//			break
	//		}
	//	}
	//}
}

func TestSpeed(t *testing.T) {
	result := testing.Benchmark(BenchmarkSkipList)
	fmt.Printf("SkipList Search Benchmark: %s\n", result)

	result = testing.Benchmark(BenchmarkStdList)
	fmt.Printf("StdList Search Benchmark: %s\n", result)

}

// 向跳表中插入测试数据
func insertSkipList(sl *list.SkipList) {
	for i := 0; i < testSize; i++ {
		score := rand.Float64() * 1000
		sl.Insert(fmt.Sprintf("Item-%d", i), int(score))
	}
}

// 向 Go 标准库的链表中插入测试数据
func insertStdList(stdList *cl.List) {
	for i := 0; i < testSize; i++ {
		score := rand.Float64() * 1000
		stdList.PushBack(int(score))
	}
}
