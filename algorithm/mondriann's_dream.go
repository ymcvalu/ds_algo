package algorithm

// 求把N*M的棋盘分割成若干个1*2的的长方形，有多少种方案
// 例如当N=2，M=4时，共有5种方案。当N=2，M=3时，共有3种方案
// |----|----| |----|--|--| |--|--|----| |--|----|--| |--|--|--|--|
// |----|----| |----|  |  | |  |  |----| |  |----|  | |  |  |  |  |
// |----|----| |----|--|--| |--|--|----| |--|----|--| |--|--|--|--|
// 可以看到，可以横着进行分割1x2，也可竖着分割2x1
//
// 数据范围：1 <= N, M <= 11

// 思路：先放横着的，横着的确定了，竖着的自然就确定了，只要每一列内部未使用的空格有连续的偶数个
// 总方案数就等于横着放的合法方案数，这里的合法是指未使用的空格可以用竖的填充
// 使用状态压缩DP：
// 			状态表示：f(i,j)表示前i-1列已经摆放完整，且使用第i-1列和第i列分割的状态j的方案数，这里状态j的每一位表示第i列的每一行是否已经使用
// 			状态计算：f(i,j) = ∑f(i-1,k), k是所有能转移到j的状态

func dreamOfMondriann(n, m int) int {
	if n <= 0 || m <= 0 {
		return -1
	}

	ss := 1 << uint(n)

	// 预处理，用于快速判断f(i,j)中的j是否有效
	st := map[int]bool{}

	for i := 0; i < ss; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			// 如果第j行不是空格，判断前面有没有连续偶数个空格
			if (i>>uint(j))&1 > 0 {
				// 如果没有连续的偶数个空格
				if cnt&1 > 0 {
					break
				}
				cnt = 0
			} else {
				cnt++
			}
		}

		st[i] = cnt&1 == 0
	}

	// 记录状态之间的转移关系
	state := make([][]int, ss)

	// 计算能转移到状态i的所有有效状态j
	//	 1. 行不能有重叠：i&j==0
	//   2. 转移之后状态j的空格要是连续偶数：st[i|j]==true
	for i := 0; i < ss; i++ {
		for j := 0; j < ss; j++ {
			if (i&j) == 0 && st[i|j] {
				state[i] = append(state[i], j)
			}
		}
	}

	// f(i,j)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, ss)
	}

	// 初始状态
	dp[0][0] = 1
	// 从第2列开始
	for i := 1; i <= m; i++ {
		// 遍历每一种状态
		for j := 0; j < ss; j++ {
			// f(i,j) = ∑f(i-1,k), k是所有能转移到j的状态
			for _, k := range state[j] {
				dp[i][j] += dp[i-1][k]
			}
		}
	}

	return dp[m][0]
}
