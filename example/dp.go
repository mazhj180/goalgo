package example

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 暴力穷举 preIdx 前一个数字的索引 currentIdx 当前数字的索引
func longestIncreasingSubsequenceHelper(items []int, preIdx, currentIdx int) (l int) {
	if currentIdx == len(items) {
		return 0
	}
	token := 0
	if preIdx == -1 || items[preIdx] < items[currentIdx] {
		token = 1 + longestIncreasingSubsequenceHelper(items, currentIdx, currentIdx+1)
	}
	notToken := longestIncreasingSubsequenceHelper(items, preIdx, currentIdx+1)

	return max(token, notToken)
}

func LongestIncreasingSubsequence(items []int) int {
	return longestIncreasingSubsequenceHelper(items, -1, 0)
}

// 动态规划
// todo

// ------------------------------------------------------------------------------
//一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 10 级的台阶总共有多少种跳法。

// FrogJumpingStepsHelper 暴力穷举
func FrogJumpingStepsHelper(n int) int {
	if n < 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return FrogJumpingStepsHelper(n-1) + FrogJumpingStepsHelper(n-2)
}
