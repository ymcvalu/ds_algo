package algorithm

func bubbleSort(arr []int) {

	n := len(arr)
	if n < 2 {
		return
	}

	for i := n - 1; i > 0; i-- {
		swap := false
		for j := 1; j <= i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				swap = true
			}
		}
		if !swap {
			break
		}
	}

}
