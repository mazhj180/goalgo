package test

import (
	"fmt"
	"mazhj.com/goalgo/data_structures/list"
	"testing"
)

func TestCircularList(t *testing.T) {

	cl := list.NewCircularList[int]()
	for _, v := range data {
		cl.Insert(v)
	}

	cl.Traverse(func(v int) {
		fmt.Printf("%d,", v)
	})
}
