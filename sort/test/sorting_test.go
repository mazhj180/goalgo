package test

import (
	"fmt"
	"mazhj.com/goalgo/sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println(randomData)
	sort.BubbleSort(randomData)
	fmt.Println(randomData)
}

func TestHeapSort(t *testing.T) {
	fmt.Println(randomData)
	sort.HeapSort(randomData)
	fmt.Println(randomData)
}

func TestMergeSort(t *testing.T) {
	fmt.Println(randomData)
	sort.MergeSort(randomData)
	fmt.Println(randomData)
}

func TestQuickSort(t *testing.T) {
	fmt.Println(randomData)
	sort.QuickSort(randomData)
	fmt.Println(randomData)
}

func TestSelectSort(t *testing.T) {
	fmt.Println(randomData)
	sort.SelectionSort(randomData)
	fmt.Println(randomData)
}

func TestInsertionSort(t *testing.T) {
	fmt.Println(randomData)
	sort.InsertSort(randomData)
	fmt.Println(randomData)
}
