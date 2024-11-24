package sort

import "cmp"

func BubbleSort[T cmp.Ordered](arr []T) {

	n, cnt := len(arr), 0

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				cnt++
			}
		}
		if cnt == 0 {
			return
		}
	}

}
