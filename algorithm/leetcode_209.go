package algorithm

/**
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ret := 0
	l := 0
	sum := 0

	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for l < r && sum-nums[l] >= s {
			sum -= nums[l]
			l++
		}

		if sum >= s && (ret == 0 || r-l+1 < ret) {
			ret = r - l + 1
			if ret == 1 {
				break
			}
		}
	}

	return ret
}
