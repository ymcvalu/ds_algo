package algorithm

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := nums[0]
	max := dp
	for i := 1; i < len(nums); i++ {
		if nums[i]+dp > nums[i] {
			dp += nums[i]
		} else {
			dp = nums[i]
		}
		if dp > max {
			max = dp
		}
	}
	return max
}
