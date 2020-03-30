package algorithm

import "sort"

/**
给出一个由无重复的正整数组成的集合，找出其中最大的整除子集，子集中任意一对 (Si，Sj) 都要满足：Si % Sj = 0 或 Sj % Si = 0。

如果有多个目标子集，返回其中任何一个均可。

*/

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	dp[0] = append(dp[0], nums[0])
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if len(dp[j])+1 > len(dp[i]) {
					dp[i] = append(dp[j], nums[i])
				}
			}
		}
		if len(dp[i]) == 0 {
			dp[i] = []int{nums[i]}
		}
	}

	m := 0
	for i := 1; i < len(nums); i++ {
		if len(dp[i]) > len(dp[m]) {
			m = i
		}
	}

	return dp[m]

}
