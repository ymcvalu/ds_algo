package algorithm

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按照区间起始进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]
	})

	result := make([][]int, 0)
	interval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if interval[1] >= intervals[i][0] {
			end := intervals[i][1]
			if interval[1] > end {
				end = interval[1]
			}
			interval = []int{interval[0], end}
		} else {
			result = append(result, interval)
			interval = intervals[i]
		}
	}
	result = append(result, interval)
	return result
}
