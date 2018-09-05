package algorithm

/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.
Example 1:
	nums1 = [1, 3]
	nums2 = [2]

	The median is 2.0

Example 2:
	nums1 = [1, 2]
	nums2 = [3, 4]

	The median is (2 + 3)/2 = 2.5
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	total := l1 + l2
	mid := total / 2
	if total%2 == 0 {
		return float64(getKth(nums1, nums2, mid)+getKth(nums1, nums2, mid+1)) / 2.0
	} else {
		return float64(getKth(nums1, nums2, mid+1))
	}
}

func getKth(nums1, nums2 []int, kth int) int {
	l1 := len(nums1)
	l2 := len(nums2)

	if l1 > l2 {
		nums1, nums2 = nums2, nums1
		l1, l2 = l2, l1
	}

	if l1 == 0 {
		return nums2[kth-1]
	}

	if kth == 1 {
		val := nums1[0]
		if val > nums2[0] {
			val = nums2[0]
		}
		return val
	}

	p := kth / 2
	if p > l1 {
		p = l1
	}

	q := kth - p
	if nums1[p-1] >= nums2[q-1] && (q == l2 || nums1[p-1] <= nums2[q]) {
		return nums1[p-1]
	} else if nums1[p-1] < nums2[q-1] {
		return getKth(nums1[p:], nums2, kth-p)
	} else {
		return getKth(nums1[:p-1], nums2, kth)
	}

}
