package algorithm


/**
输入一个整形数组，数组里面有正数也有负数。数组中一个或者连续多个整数组成一个子数组。求所有子数组中和的最大值。要求时间复杂度为O(n)
 */
func findGreatestSumOfSubarray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	dp := make([]int, len(arr))
	dp[0] = arr[0]
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if sum := dp[i-1] + arr[i]; sum > arr[i] {
			dp[i] = sum
			if dp[i] > max {
				max = dp[i]
			}
		} else {
			dp[i] = arr[i]
		}
	}

	return max
}
