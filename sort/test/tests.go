package test

import (
	"fmt"
	"time"
)

// test case
var (
	// 随机数据
	randomData = []int{47, 3, 76, 23, 8, 65, 15, 95, 34, 58}

	// 有序数据
	sortedData = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// 逆序数据
	reversedData = []int{99, 85, 76, 64, 52, 41, 33, 22, 15, 8}

	// 重复数据
	repeatedData = []int{42, 42, 42, 42, 42, 42, 42, 42, 42, 42}

	// 几乎有序数据
	nearlySortedData = []int{1, 3, 5, 7, 8, 10, 12, 11, 15, 14}
)

func spent(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Printf("执行耗时 : %s", elapsed)
}
