package algorithm

/**
给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
*/

func waysToChange(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	costs := []int{1, 5, 10, 25}
	for _, c := range costs {
		for i := c; i <= n; i++ {
			dp[i] += dp[i-c]
		}
	}
	return dp[n] % 1000000007
}
