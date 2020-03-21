package algorithm

import "sort"

/**
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
1,3,2 → 2,1,3
*/

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

	i := len(nums) - 1
	for i > 0 {
		if nums[i] > nums[i-1] {
			break
		}
		i--
	}

	if i > 0 {
		v := nums[i-1]
		for j := len(nums) - 1; j >= i; j-- {
			if nums[j] > v {
				nums[j], nums[i-1] = nums[i-1], nums[j]
				sort.Ints(nums[i:])

				break
			}
		}
	} else {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
