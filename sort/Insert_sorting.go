package sort

func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > arr[i] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}
