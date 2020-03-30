package algorithm

import "container/list"

/**
你现在手里有一份大小为 N x N 的「地图」（网格） grid，上面的每个「区域」（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，请你找出一个海洋区域，这个海洋区域到离它最近的陆地区域的距离是最大的。

我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 - x1| + |y0 - y1| 。

如果我们的地图上只有陆地或者海洋，请返回 -1。
*/

// BFS同时对所有陆地进行搜索，最后一个遍历的节点就是最远的海洋了
func maxDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	type Coordinate struct {
		x int
		y int
		d int
	}

	q := list.New()
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				q.PushBack(Coordinate{i, j, 0})
			}
		}
	}

	if q.Len() == len(grid)*len(grid) || q.Len() == 0 {
		return -1
	}

	for {
		elem := q.Front()
		c := elem.Value.(Coordinate)
		q.Remove(elem)

		if c.x > 0 && grid[c.x-1][c.y] == 0 {
			grid[c.x-1][c.y] = 1
			q.PushBack(Coordinate{c.x - 1, c.y, c.d + 1})
		}

		if c.x+1 < len(grid) && grid[c.x+1][c.y] == 0 {
			grid[c.x+1][c.y] = 1
			q.PushBack(Coordinate{c.x + 1, c.y, c.d + 1})
		}

		if c.y > 0 && grid[c.x][c.y-1] == 0 {
			grid[c.x][c.y-1] = 1
			q.PushBack(Coordinate{c.x, c.y - 1, c.d + 1})
		}

		if c.y+1 < len(grid) && grid[c.x][c.y+1] == 0 {
			grid[c.x][c.y+1] = 1
			q.PushBack(Coordinate{c.x, c.y + 1, c.d + 1})
		}

		if q.Len() == 1 {
			break
		}
	}

	return q.Front().Value.(Coordinate).d
}
