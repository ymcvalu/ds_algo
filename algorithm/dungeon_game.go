package algorithm

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	li := len(dungeon)
	lj := len(dungeon[0])
	dp := make([][]int, li)
	for i := 0; i < li; i++ {
		dp[i] = make([]int, lj)
	}

	dp[li-1][lj-1] = 1

	for i := li - 1; i >= 0; i-- {
		for j := lj - 1; j >= 0; j-- {
			if i == li-1 {
				if j == lj-1 {
					continue
				}
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j+1])
			} else if j == lj-1 {
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i+1][j])
			} else {
				dp[i][j] = min(max(1, dp[i+1][j]-dungeon[i+1][j]), max(1, dp[i][j+1]-dungeon[i][j+1]))
			}
		}
	}

	return max(1, dp[0][0]-dungeon[0][0])
}
