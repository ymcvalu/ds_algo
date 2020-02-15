package algorithm

// 割绳子：一段长度为k的绳子，将其分割成多段（至少两段），然后将每段长乘积，求最大值
// 比如 2=> 1*1
//      6=> 3*3
func splitRopeDP(l int) int {
	if l < 1 {
		panic("invalid length")
	}

	if l == 1 {
		return 1
	}

	dp := make([]int, l+1)
	dp[1] = 1

	for i := 2; i <= l; i++ {
		m := 0
		for j := 1; j < i; j++ {
			m = max(m, j*dp[i-j])
		}
		if i < l {
			m = max(m, i)
		}
		dp[i] = m
	}

	return dp[l]
}
