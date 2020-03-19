package algorithm

/**
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

说明：

不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。
*/

func findDuplicate(nums []int) int {
	tortoise := nums[0]
	hare := nums[0]

	// 可以把val当做是指向下一个节点的指针，如果数组中存在相同的元素，说明存在环
	// 因为数字在1到n之间，因此遍历的时候不会经过num[0]
	// 快慢指针检测环
	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	ptr1 := tortoise
	ptr2 := nums[0]
	// hare的速度是tortoise的两倍，因此ptr1和ptr2最终会在环的入口相遇，也就是重复的那个数
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}
	return ptr1
}
