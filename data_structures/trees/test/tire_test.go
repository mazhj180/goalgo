package test

import (
	"fmt"
	"goalgo/data_structures/trees"
	"testing"
)

var data = []string{"你好", "你好呀", "你好吗", "你是谁", "我不在乎", "我不在意", "我和你"}

func TestTire(t *testing.T) {
	tire := trees.Tire{}
	for _, v := range data {
		tire.AddWord(v)
	}
	for _, v := range tire.Retrieve("你好") {
		fmt.Println(v)
	}
}
