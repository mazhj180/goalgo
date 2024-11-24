package sort

import "cmp"

func MergeSort[T cmp.Ordered](arr []T) {

	var dfs func(left, right int) []T

	dfs = func(left, right int) []T {
		if left == right {
			return arr[left : right+1]
		}

		mid := left + (right-left)/2

		arr1 := dfs(left, mid)
		arr2 := dfs(mid+1, right)

		return merge(arr1, arr2)
	}

	sorted := dfs(0, len(arr)-1)
	copy(arr, sorted)
}

func merge[T cmp.Ordered](arr1, arr2 []T) []T {

	len1, len2 := len(arr1), len(arr2)
	merged := make([]T, 0, len(arr1)+len(arr2))

	i, j := 0, 0
	for i < len1 && j < len2 {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	merged = append(merged, arr1[i:]...)
	merged = append(merged, arr2[j:]...)

	return merged
}
