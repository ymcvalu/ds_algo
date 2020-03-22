package algorithm

import "sort"

/**
给定整数数组 A，每次 move 操作将会选择任意 A[i]，并将其递增 1。
返回使 A 中的每个值都是唯一的最少操作次数。
*/

// 方法1：先排序，然后如果当前元素小于或者等于前一个元素，则增长为前一个元素+1，时间复杂度为O(n log n)
func minIncrementForUnique1(A []int) int {
	times := 0
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			pre := A[i]
			A[i] = A[i-1] + 1
			times += A[i] - pre
		}
	}
	return times
}

// 方法2：空间换时间，使用数组记录每个元素出现的次数，如果某个元素出现超过一次，则增加到当前最小的没有出现过的元素，这里寻找需要一定的技巧：先记录，等下一次遇到出现次数为0的元素再处理，总的时间复杂度为O(n)：
func minIncrementForUnique2(A []int) int {
	arr := [40001]int{} // 题目限制元素最大是40000，因此直接使用数组，否则需要使用map或者先找出当前最大元素
	// 记录每个元素出现次数
	for i := 0; i < len(A); i++ {
		arr[A[i]] += 1
	}
	times := 0 // 记录替换次数
	t := 0
	i := 0
	nums := 0 // 记录处理过的元素个数
	for ; nums < len(A); i++ {
		if arr[i] == 0 {
			if t > 0 {
				t--
				times += i
			}
			continue
		}

		nums += arr[i] // 记录处理的元素个数
		for arr[i] > 1 {
			arr[i]--
			t++
			times -= i
		}
	}
	for t > 0 {
		times += i
		i++
		t--
	}
	return times
}
