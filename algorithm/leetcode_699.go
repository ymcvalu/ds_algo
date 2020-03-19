package algorithm

/**
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

注意: 给定的矩阵grid 的长度和宽度都不超过 50。
*/

func maxAreaOfIsland(grid [][]int) int {
	max := 0

	type P struct {
		i, j int
	}
	xy := [4][4]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			grid[i][j] = 0
			stk := []P{{i, j}}
			t := 1
			c := 0
			for t > 0 {
				p := stk[t-1]
				t--
				stk = stk[:t]
				c++

				for _, _xy := range xy {
					nx := p.i + _xy[0]
					if nx < 0 || nx >= len(grid) {
						continue
					}
					ny := p.j + _xy[1]
					if ny < 0 || ny >= len(grid[0]) {
						continue
					}
					if grid[nx][ny] == 0 {
						continue
					}
					grid[nx][ny] = 0
					stk = append(stk, P{nx, ny})
					t++
				}

			}

			if c > max {
				max = c
			}

		}
	}

	return max

}
