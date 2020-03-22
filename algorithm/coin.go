package algorithm

import "sort"

/**
假设有n个不同面值的零钱（最小的面值是1元，并且肯定存在），个数无限，现在想凑N元，怎么凑才能使钱的个数最少
*/

func coin(n int, coins []int) (int, map[int]int) {
	if n == 0 || len(coins) == 0 {
		return -1, nil
	}

	// 零钱有序排列
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})

	// 最小值要为1
	if coins[0] != 1 {
		return -1, nil
	}

	dp := make([]int, n+1)

	// dp求最值
	for i := 1; i <= n; i++ {
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			if dp[i] == 0 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	solve := map[int]int{}
	c := dp[n]
	m := n
	// 回溯，求找零方案
	for i := n - 1; i >= 0; i-- {
		if dp[i] == c-1 {
			solve[m-i] += 1
			c = dp[i]
			m = i
		}
	}

	return dp[n], solve
}
