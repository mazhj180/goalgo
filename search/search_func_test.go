package search

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 80}
	idx := SeqSearch(arr, 80)
	fmt.Println(idx)
}
