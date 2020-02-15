package algorithm

// 求数组逆序对数量
func InversePairs(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	return _inversePairs(arr)

}

// 可以在归并排序过程中计算逆序对数量
// arr[i,j] => left:arr[i,mid) || right:arr[mid,j]
// 现在假设left和right都已经是有序的数组了
// 在归并过程中，如果存在right中的元素小于left中的元素，则记录逆序对数量
func _inversePairs(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	narr := make([]int, len(arr))
	copy(narr, arr)
	mid := len(narr) / 2

	left := narr[:mid]
	right := narr[mid:]

	count := _inversePairs(left) + _inversePairs(right)
	lc, rc := len(left), len(right)
	li, ri := 0, 0
	arr = arr[:0]
	for li < lc && ri < rc {
		if left[li] < right[ri] {
			arr = append(arr, left[li])
			li++
		} else if right[ri] < left[li] {
			// 记录逆序对数量
			count += lc - li
			arr = append(arr, right[ri])
			ri++
		} else {
			arr = append(arr, left[li])
			li++
			arr = append(arr, right[ri])
			ri++
		}
	}
	for li < lc {
		arr = append(arr, left[li])
		li++
	}
	for ri < rc {
		arr = append(arr, right[ri])
		ri++
	}

	return count
}
