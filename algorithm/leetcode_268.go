package algorithm

/**
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

示例 1:

输入: [3,0,1]
输出: 2

示例 2:

输入: [9,6,4,2,3,5,7,0,1]
输出: 8
*/
func missingNumber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	hasOne := false
	hasZero := false
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			hasOne = true
		}
		if nums[i] == 0 {
			hasZero = true
			nums[i] = 1
		}
	}
	if !hasZero {
		return 0
	}
	if !hasOne {
		return 1
	}

	for i := 0; i < n; i++ {
		c := nums[i]

		if c < 0 {
			c = -c
		}
		if c == n {
			c = 0
		}

		if nums[c] > 0 {
			nums[c] *= -1
		}
	}

	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return n
	}
	return n + 1
}
