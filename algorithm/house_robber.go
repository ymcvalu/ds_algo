package algorithm

/**

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:
	Input: [1,2,3,1]
	Output: 4
	Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3). Total amount you can rob = 1 + 3 = 4.

Example 2:
	Input: [2,7,9,3,1]
	Output: 12
	Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1). Total amount you can rob = 2 + 9 + 1 = 12.

*/

func rob(nums []int) ([]int, int) {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	ln := len(nums)
	if ln == 0 {
		return nil, 0
	} else if ln == 1 {
		return []int{0}, nums[0]
	}

	var ms = make([]int, len(nums))
	ms[0] = nums[0]
	ms[1] = max(nums[0], nums[1])
	for i := 2; i < ln; i++ {
		ms[i] = max(ms[i-1], ms[i-2]+nums[i])
	}

	path := make([]int, 0)

	n := ms[ln-1]
	for i := ln - 1; i > 1; i-- {
		if ms[i] == n && n == ms[i-2]+nums[i] {
			path = append(path, i)
			n -= nums[i]
		}
	}

	if n > 0 {
		if n == nums[0] {
			path = append(path, 0)
		} else {
			path = append(path, 1)
		}
	}

	return path, ms[ln-1]
}
