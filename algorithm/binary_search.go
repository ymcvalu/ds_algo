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
