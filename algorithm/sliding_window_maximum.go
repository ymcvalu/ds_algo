package algorithm

import "container/list"

/**
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window.
Each time the sliding window moves right by one position.
Return the max sliding window

Example:
Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7

*/

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k == 0 {
		return nil
	}

	rets := make([]int, 0, len(nums)-k+1)
	deque := list.New()
	for i := 0; i < len(nums); i++ {
		for deque.Len() > 0 && nums[deque.Back().Value.(int)] < nums[i] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)
		if deque.Front().Value.(int)+k == i {
			deque.Remove(deque.Front())
		}

		if i >= k-1 {
			rets = append(rets, nums[deque.Front().Value.(int)])
		}
	}
	return rets
}
