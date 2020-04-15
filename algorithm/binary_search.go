package algorithm

func binarySearch(arr []int, v int) int  {
	if len(arr) == 0 {
		return -1
	}

	p := 0
	q := len(arr) - 1
	for p <= q {
		mid := (p + q) >>1
		if arr[mid] == v {
			return mid
		}

		if arr[mid] < v {
			p = mid + 1
		} else {
			q = mid - 1
		}

	}

	return -1
}

func binary(arr []int, n int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) >> 1
		if arr[mid] < n {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
