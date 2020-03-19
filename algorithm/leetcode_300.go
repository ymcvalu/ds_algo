package algorithm

/**
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
*/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				l := dp[j] + 1
				if l > dp[i] {
					dp[i] = l
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
