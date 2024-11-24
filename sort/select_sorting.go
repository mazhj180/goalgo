package sort

import "cmp"

func SelectionSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 假设当前元素是最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换当前元素和找到的最小元素
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
