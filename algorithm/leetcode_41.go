package algorithm

/**
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3

示例 2:

输入: [3,4,-1,1]
输出: 2

说明:

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

*/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	hasOne := false
	// 将大于n的或者小于等于0的数替换为1，同时检查数组中是否存在1
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			hasOne = true
		}
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	if !hasOne {
		return 1
	}

	// 现在数组中，存在的元素都是1~n
	// 如果i存在，则将nums[i/mod n]的符号为置为负数，表示i存在
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
