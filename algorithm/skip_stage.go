package algorithm

// 跳台阶，每次可以跳一级台阶或者跳两级台阶，跳到第n级台阶共有几种方法
// 思路：
// 		dp: fn(n) = fn(n-1) + f(n-2)
// 实际上就是斐波那契数列

func skipStage(n int) int {
	if n == 0 {
		return 0
	}
	dp1 := 1
	dp2 := 1
	for i := 2; i <= n; i++ {
		dp2, dp1 = dp2+dp1, dp2
	}
	return dp2
}
