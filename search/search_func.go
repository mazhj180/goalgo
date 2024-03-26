package search

// SeqSearch 顺序查找 O(N)
func SeqSearch(arr []int, tar int) int {
	for idx, val := range arr {
		if val == tar {
			return idx
		}
	}
	return -1
}

// BinarySearch 二分查找 O(logN)
func BinarySearch(arr []int, tar int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		midIdx := (left + right) >> 1
		if tar < arr[midIdx] {
			right = midIdx - 1
		} else if tar > arr[midIdx] {
			left = midIdx + 1
		} else {
			var indices []int
			cur := midIdx - 1
			for {
				if cur < 0 || arr[cur] != tar {
					break
				}
				indices = append(indices, cur)
				cur -= 1
			}
			indices = append(indices, midIdx)
			cur = midIdx + 1
			for {
				if cur > len(arr)-1 || arr[cur] != tar {
					break
				}
				indices = append(indices, cur)
				cur += 1
			}
			return indices
		}
	}
	return []int{}
}

// BinarySearchR 二分查找 (递归)
func BinarySearchR(arr []int, left, right, tar int) []int {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	if tar < arr[mid] {
		return BinarySearchR(arr, left, mid-1, tar)
	} else if tar > arr[mid] {
		return BinarySearchR(arr, mid+1, right, tar)
	} else {
		var indices []int
		cur := mid - 1
		for {
			if cur < 0 || cur != tar {
				break
			}
			indices = append(indices, cur)
			cur -= 1
		}
		indices = append(indices, mid)
		cur = mid + 1
		for {
			if cur > len(arr)-1 || cur != tar {
				break
			}
			indices = append(indices, cur)
			mid += 1
		}
		return indices
	}
}

// 二分查找(回顾)
func binarySearchP(arr []int, target int) []int {
	left := 0
	right := len(arr)
	for left <= right {
		mid := (left + right) >> 1
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			var indices []int
			cur := mid - 1
			for {
				if cur < 0 || arr[cur] != target {
					break
				}
				indices = append(indices, cur)
				cur -= 1
			}
			indices = append(indices, cur)
			for {
				if cur > len(arr) || arr[cur] != target {
					break
				}
				indices = append(indices, cur)
				cur += 1
			}
			return indices
		}
	}
	return []int{}
}
