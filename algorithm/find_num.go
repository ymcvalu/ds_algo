package algorithm

func findNum(arr []int, num int) bool {
	if len(arr) == 0 {
		return false
	} else if len(arr) == 1 {
		return arr[0] == num
	}
	maxIdx := findMax(arr, 0)
	return binSearch(arr[:maxIdx], num, true) || binSearch(arr[maxIdx:], num, false)
}

func binSearch(arr []int, num int, asc bool) bool {
	if len(arr) == 0 {
		return false
	}
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) >> 1
		if arr[mid] == num {
			return true
		}

		if arr[mid] > num {
			if asc {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if asc {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return arr[l] == num
}

func findMax(arr []int, i int) int {
	if len(arr) == 1 {
		return i
	}

	mid := (len(arr) - 1) / 2

	if mid+1 == len(arr) || arr[mid+1] > arr[mid] {
		return findMax(arr[mid+1:], i+mid+1)
	} else {
		return findMax(arr[:mid+1], i)
	}
}
