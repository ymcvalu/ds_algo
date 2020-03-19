package algorithm

/**
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
*/

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func countBitsDP(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			dp[i] = dp[i>>1]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}

	return dp
}
