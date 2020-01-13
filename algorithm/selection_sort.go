package algorithm

func SelectionSort(arr []int) {
	ln := len(arr)
	if ln < 2 {
		return
	}

	for i := 0; i < ln; i++ {
		min := i
		for j := i + 1; j < ln; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

}
