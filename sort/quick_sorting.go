package sort

import "cmp"

func QuickSort[T cmp.Ordered](arr []T) {

	var dfs func(left, right int)

	dfs = func(left, right int) {
		if left > right {
			return
		}

		// 快慢指针划分左右区间
		fast, slow := left, left

		for fast < right {
			if arr[fast] < arr[right] {
				arr[slow], arr[fast] = arr[fast], arr[slow]
				slow++
			}
			fast++
		}

		arr[slow], arr[fast] = arr[fast], arr[slow]

		//递归左右区间
		dfs(left, slow-1)
		dfs(slow+1, right)
	}

	dfs(0, len(arr)-1)
}
