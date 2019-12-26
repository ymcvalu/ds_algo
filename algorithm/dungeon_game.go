package algorithm

import "unsafe"

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}

	// max int
	min := (1 << (8*unsafe.Sizeof(int(1)) - 1)) - 1
	calcMinHP(dungeon, 0, 0, 0, 0, &min)
	return min
}

func calcMinHP(dungeon [][]int, i, j int, max, h int, curMin *int) {
	h += dungeon[i][j]
	if h < max {
		max = h
	}

	if i == len(dungeon)-1 && j == len(dungeon[0])-1 {
		if max == 0 {
			max = 1
		} else if max < 0 {
			max = 1 - max
		}
		if max < *curMin {
			*curMin = max
		}
	}

	// 减枝
	if 1-max > *curMin {
		return
	}

	if i+1 < len(dungeon) {
		calcMinHP(dungeon, i+1, j, max, h, curMin)
	}
	if j+1 < len(dungeon[0]) {
		calcMinHP(dungeon, i, j+1, max, h, curMin)
	}
}
