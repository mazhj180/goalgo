package kmp_test

import (
	"mazhj.com/goalgo/kmp"
	"testing"
)

func TestKMP(t *testing.T) {
	needle := "ababc"
	haystack := "abababc"

	// 构建 next 表
	next := kmp.BuildNextTable(needle)

	// 匹配测试
	result := next.MatchStr(haystack, needle)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}

	// 测试不匹配情况
	result = next.MatchStr("abcdef", "gh")
	if result != -1 {
		t.Errorf("Expected -1, but got %d", result)
	}

	// 测试部分匹配情况
	result = next.MatchStr("ababcabc", "abc")
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}
