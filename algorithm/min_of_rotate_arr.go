package algorithm

// 把一个升序数列在某个位置进行旋转，导致前面若干个元素被移动到数组末尾，输出数组的最小值
// 1 2 3 4 5
// 3 4 5 1 2
func MinOfRotateArr(arr []int) int {
	if len(arr) == 0 {
		panic("invalid param")
	}
	l, r := 0, len(arr)-1
	if l == r || arr[l] < arr[r] {
		return arr[l]
	}

	for l < r {
		mid := (l + r) >> 1
		if arr[mid] > arr[len(arr)-1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return arr[r]

}
