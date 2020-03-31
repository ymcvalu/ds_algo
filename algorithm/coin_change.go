package algorithm

import "sort"

/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
*/

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = -1
	}
	for i := coins[0]; i <= amount; i++ {
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			if dp[i-coins[j]] >= 0 {
				if dp[i] == -1 {
					dp[i] = dp[i-coins[j]] + 1
				} else {
					dp[i] = min(dp[i], dp[i-coins[j]]+1)
				}
			}
		}
	}

	return dp[amount]
}
