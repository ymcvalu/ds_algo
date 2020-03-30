package algorithm

import "unsafe"

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
*/

func singleNumber(nums []int) int {
	n := 0
	size := uint(8 * unsafe.Sizeof(1))
	for s := uint(0); s < size; s++ {
		c := 0
		for i := 0; i < len(nums); i++ {
			if (nums[i]>>s)&1 == 1 {
				c += 1
			}
		}
		if c%3 == 1 {
			n |= 1 << s
		}
	}
	return n
}
