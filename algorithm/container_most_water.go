package algorithm

/**
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

func maxArea(height []int) int {
	var (
		i   = 0
		j   = len(height) - 1
		max = 0
	)
	if i >= j {
		return 0
	}

	for i < j {
		area := min(height[i], height[j]) * (j - i)
		if area > max {
			max = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
