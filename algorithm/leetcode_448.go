package algorithm

/**
给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。

找到所有在 [1, n] 范围之间没有出现在数组中的数字。

您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
*/

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
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

	ret := make([]int, 0, len(nums))
	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			ret = append(ret, i)
		}
	}
	if nums[0] > 0 {
		ret = append(ret, n)
	}
	return ret
}
