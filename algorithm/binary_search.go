package algorithm

func binarySearch(arr []int, v int) bool {
	if len(arr) == 0 {
		return false
	}

	p := 0
	q := len(arr) - 1
	for p <= q {
		mid := (p + q) / 2
		if arr[mid] == v {
			return true
		}

		if arr[mid] < v {
			p = mid + 1
		} else {
			q = mid - 1
		}

	}

	return false
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
