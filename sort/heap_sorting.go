package sort

import "cmp"

func HeapSort[T cmp.Ordered](arr []T) {

	n := len(arr)

	// 数组堆化
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 排序
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

}

func heapify[T cmp.Ordered](arr []T, n, k int) {

	for k < n>>1 {

		largest := 2*k + 1
		right := 2*k + 2

		if right < n && arr[right] > arr[largest] {
			largest = right
		}

		if arr[largest] <= arr[k] {
			break
		}

		arr[k], arr[largest] = arr[largest], arr[k]

		k = largest
	}

}
