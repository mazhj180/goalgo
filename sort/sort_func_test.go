package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

// 打印
func display(str string, arr []int) {
	fmt.Println(str)
	for _, val := range arr {
		fmt.Printf("%d\t", val)
	}
	fmt.Printf("\n")
}

// 初始化切片
func initSli(length, ran int) {
	arr = make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(ran)
	}
}

var arr []int

func Test(t *testing.T) {
	initSli(100, 1000)
	display("排序前", arr)
	bubbleSort(arr)
	display("排序后", arr)
}
