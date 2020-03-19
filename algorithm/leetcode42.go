package algorithm

// https://leetcode-cn.com/problems/trapping-rain-water/
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 示例:
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出: 6
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	right := make([]int, len(height))
	right[n-1] = height[n-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(height[i], right[i+1])
	}

	c := 0
	l := height[0]
	for i := 0; i < n; i++ {
		if height[i] > l {
			l = height[i]
		}

		h := min(l, right[i])
		c += h - height[i]
	}
	return c
}
