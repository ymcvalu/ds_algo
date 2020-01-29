package algorithm

// 把一个升序数列在某个位置进行旋转，导致前面若干个元素被移动到数组末尾，输出数组的最小值
func MinOfRotateArr(arr []int) int {
	if len(arr) == 0 {
		panic("invalid param")
	}
	l, r := 0, len(arr)-1
	if l == r || arr[l] < arr[r] {
		return arr[l]
	}

	for arr[l] >= arr[r] {
		if r-l == 1 {
			break
		}
		mid := (l + r) / 2
		if arr[l] <= arr[mid] {
			l = mid
		} else {
			r = mid
		}
	}
	return arr[r]
}
