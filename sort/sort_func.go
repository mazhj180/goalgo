package sort

import (
	"math"
)

// 冒泡排序 最差时间复杂度 O(n^2)
func bubbleSort(arr []int) {
	temp, count := 0, 0
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				count++
			}
		}
		if count == 0 {
			break
		}
	}
}

// 选择排序  最差时间复杂度 O(n^2)
func selectSort(arr []int) {
	var temp int
	for i := 0; i < len(arr); i++ {
		minVal := arr[i]
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < minVal {
				minVal = arr[j]
				minIdx = j
			}
		}
		if minIdx != i {
			temp = arr[minIdx]
			arr[minIdx] = arr[i]
			arr[i] = temp
		}
	}
}

// 插入排序 最坏时间复杂度 O(n^2)
func inSort(arr []int) {
	var in int
	for i := 1; i < len(arr); i++ {
		in = arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > in {
				arr[j+1] = arr[j]
				arr[j] = in
			}
		}
	}
}

// 希尔排序 exchange
func shellSortExchange(arr []int) {
	var temp int
	for gap := len(arr) >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < len(arr); i++ {
			for j := i - gap; j >= 0; j -= gap {
				if arr[j] > arr[j+gap] {
					temp = arr[j]
					arr[j] = arr[j+gap]
					arr[j+gap] = temp
				}
			}
		}
	}
}

// 希尔排序 插入 平均时间复杂度 O(nlog2n)
func shellSortIn(arr []int) {
	for gap := len(arr) >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < len(arr); i++ {
			inVal := arr[i]
			for j := i - gap; j >= 0; j -= gap {
				if arr[j] > inVal {
					arr[j+gap] = arr[j]
					arr[j] = inVal
				}
			}
		}
	}

}

// 快速排序
func quickSort(arr []int, left, right int) {
	sentry := arr[(left+right)>>1]
	l := left
	r := right
	for l < r {
		for ; arr[l] < sentry; l++ {
		}
		for ; arr[r] > sentry; r-- {
		}
		if l >= r {
			break
		}
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp

		if arr[l] == sentry {
			r -= 1
		}
		if arr[r] == sentry {
			l += 1
		}
	}
	if l == r {
		l += 1
		r -= 1
	}
	if left < r {
		quickSort(arr, left, r)
	}
	if l < right {
		quickSort(arr, l, right)
	}

}

// 归并排序 迭代 O(nlogN)
func mergeSort(arr []int) {
	temp := make([]int, len(arr), cap(arr))
	length := len(arr)
	for block := 1; block < length; block <<= 1 {
		for i := 0; i < length; i += block << 1 {
			cur, idx1, idx2 := i, i, i+block
			limit1, limit2 := idx1+block, idx2+block
			if idx2 > length {
				idx2 = length
			}
			if limit2 > length {
				limit2 = length
			}
			if limit1 > length {
				limit1 = length
			}
			for idx1 < limit1 && idx2 < limit2 {
				if arr[idx1] < arr[idx2] {
					temp[cur] = arr[idx1]
					idx1++
				} else {
					temp[cur] = arr[idx2]
					idx2++
				}
				cur++
			}
			for idx1 < limit1 {
				temp[cur] = arr[idx1]
				idx1++
				cur++
			}
			for idx2 < limit2 {
				temp[cur] = arr[idx2]
				idx2++
				cur++
			}

		}
		swap := temp
		temp = arr
		arr = swap
	}
}

// 归并排序 递归
func mergeSortR(arr, temp []int, left, right int) {
	if left >= right {
		arr = temp
		return
	}
	length := right - left
	limit1 := length>>1 + left
	idx1, idx2, limit2 := left, limit1+1, right
	mergeSortR(arr, temp, idx1, limit1)
	mergeSortR(arr, temp, idx2, limit2)
	cur := left
	for idx1 <= limit1 && idx2 <= limit2 {
		if arr[idx1] < arr[idx2] {
			temp[cur] = arr[idx1]
			idx1++
		} else {
			temp[cur] = arr[idx2]
			idx2++
		}
		cur++
	}
	for idx1 <= limit1 {
		temp[cur] = arr[idx1]
		idx1++
		cur++
	}
	for idx2 <= limit2 {
		temp[cur] = arr[idx2]
		idx2++
		cur++
	}
	for i := left; i <= right; i++ {
		arr[i] = temp[i]
	}
}

// 基数排序
func cardinalitySort(arr []int) {

	length := len(arr)
	bucket := make([][]int, 10)
	mark := make([]int, length)

	for i := range bucket {
		bucket[i] = make([]int, length)
	}
	maxNum := 0
	for i := 0; i < length; i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}
	counts := 0
	for maxNum > 0 {
		maxNum /= 10
		counts++
	}

	for idx := 0; idx < counts; idx++ {
		//分散到不同的桶中
		for i := 0; i < length; i++ {
			div := int(math.Pow10(idx + 1))
			bit := (arr[i] * 10 / div) % 10
			bucket[bit][mark[bit]] = arr[i]
			mark[bit] += 1
		}
		//按桶序号 取出数据
		cur := 0
		for i := 0; i < 10; i++ {
			if mark[i] > 0 {
				for j := 0; j < mark[i]; j++ {
					arr[cur] = bucket[i][j]
					bucket[i][j] = 0
					cur++
				}
				mark[i] = 0
			}
		}

	}

}

//func main() {
//	arr := make([]int, 100)
//	for i := 0; i < 100; i++ {
//		arr[i] = rand.Intn(1000)
//	}
//	display("排序前 ：", arr)
//	//bubbleSort(arr)
//	//selectSort(arr)
//	//inSort(arr)
//	//shellSortExchange(arr)
//	//shellSortIn(arr)
//	//quickSort(arr, 0, len(arr)-1)
//	//mergeSort(arr)
//	//temp := make([]int, 100)
//	//mergeSortR(arr, temp, 0, len(temp)-1)
//	cardinalitySort(arr)
//	display("排序后 :", arr)
//
//}
