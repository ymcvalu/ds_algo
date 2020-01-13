package algorithm

func bubbleSort(arr []int) {
	ln := len(arr)
	if ln < 2 {
		return
	}

	swap := false
	for i := ln - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if arr[j-1] > arr[j] {
				swap = true
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		if !swap {
			break
		}
	}

}
