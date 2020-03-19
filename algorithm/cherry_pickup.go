package algorithm

/**
 In a N x N grid representing a field of cherries, each cell is one of three possible integers.
  - 0 means the cell is empty, so you can pass through;
  - 1 means the cell contains a cherry, that you can pick up and pass through;
  - -1 means the cell contains a thorn that blocks your way.

Your task is to collect maximum number of cherries possible by following the rules below:
 - Starting at the position (0, 0) and reaching (N-1, N-1) by moving right or down
	through valid path cells (cells with value 0 or 1);
 - After reaching (N-1, N-1), returning to (0, 0) by moving left or up through valid path cells;
 - When passing through a path cell containing a cherry, you pick it up and the cell becomes an empty cell (0);
 - If there is no valid path between (0, 0) and (N-1, N-1), then no cherries can be collected.

Example 1:

Input: grid =
[[0, 1, -1],
 [1, 0, -1],
 [1, 1,  1]]
Output: 5
Explanation:
The player started at (0, 0) and went down, down, right right to reach (2, 2).
4 cherries were picked up during this single trip, and the matrix becomes [[0,1,-1],[0,0,-1],[0,0,0]].
Then, the player went left, up, up, left to return home, picking up one more cherry.
The total number of cherries picked up is 5, and this is the maximum possible.

Note:
 - grid is an N by N 2D array, with 1 <= N <= 50.
 - Each grid[i][j] is an integer in the set {-1, 0, 1}.
 - It is guaranteed that grid[0][0] and grid[N-1][N-1] are not -1.

*/

// 只能得到局部最优解
// 最优解估计要用状态压缩dp
func cherryPickup(grid [][]int) int {
	n := len(grid)

	num := 0
	for r := 0; r < 2; r++ {
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				up, left := 0, 0
				if j > 0 {
					if grid[i][j-1] == -1 {
						up = -1
					} else {
						up = dp[i][j-1]
					}
				}
				if i > 0 {
					if grid[i-1][j] == -1 {
						left = -1
					} else {
						left = dp[i-1][j]
					}
				}

				m := max(up, left)
				if grid[i][j] == -1 || m == -1 {
					dp[i][j] = -1
				} else {
					dp[i][j] = m + grid[i][j]
				}
			}
		}

		if dp[n-1][n-1] > 0 {
			num += dp[n-1][n-1]
		}

		if r == 0 {
			i, j := n-1, n-1
			for {

				if j > 0 {
					if dp[i][j-1] == dp[i][j]-grid[i][j] {
						grid[i][j] = 0
						j--
						continue
					}
				}

				if i > 0 {
					if dp[i-1][j] == dp[i][j]-grid[i][j] {
						grid[i][j] = 0
						i--
						continue
					}
				}

				grid[i][j] = 0
				break
			}
		}

	}

	return num
}
