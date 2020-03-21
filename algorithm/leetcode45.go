package algorithm

/**
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。
*/

func jumpPriorityQueue(nums []int) int {
	pq := PriorityQueue{}
	pq.Push(0, 0)
	step := make([]int, len(nums))

loop:
	for {
		idx, ok := pq.Pop()
		if !ok {
			break
		}
		for i := 1; idx+i < len(nums) && i <= nums[idx]; i++ {
			if step[idx+1] == 0 || step[idx+1] > step[idx]+1 {
				step[idx+1] = step[idx] + 1
				if idx+1 == len(nums)-1 {
					break loop
				}
				pq.Push(idx, float64(step[idx+1]))
			}
		}

	}

	return step[len(nums)-1]
}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	steps := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i != 0 && steps[i] == 0 {
			continue
		}
		for j := 1; i+j < len(nums) && j <= nums[i]; j++ {
			if steps[i+j] == 0 || steps[i+j] > steps[i]+1 {
				steps[i+j] = steps[i] + 1
			}
		}
	}
	return steps[len(nums)-1]
}

func jumpDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		step := -1
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				if step == -1 {
					step = dp[j] + 1
				} else if step > dp[j]+1 {
					step = dp[j] + 1
				}
			}
		}
		dp[i] = step
	}
	return dp[len(nums)] - 1
}
