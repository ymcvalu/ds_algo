package algorithm

// 跳台阶，每次可以跳一阶或者跳两阶，跳到n阶共有几种方法
// 思路：
// 		dp: fn(n) = fn(n-1) + f(n+1)

func skipStage(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
