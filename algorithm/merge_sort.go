package algorithm

func mergeSort(arr []int) {
	ln := len(arr)
	if ln < 2 {
		return
	}

	mid := len(arr) / 2

	arr1 := make([]int, len(arr[:mid]))
	copy(arr1, arr[:mid])

	arr2 := make([]int, len(arr[mid:]))
	copy(arr2, arr[mid:])

	mergeSort(arr1)
	mergeSort(arr2)

	ln1 := len(arr1)
	ln2 := len(arr2)
	arr = arr[:0]
	var i, j int
	for i < ln1 && j < ln2 {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}
	if i < ln1 {
		arr = append(arr, arr1[i:]...)
	} else if j < ln2 {
		arr = append(arr, arr2[j:]...)
	}

}
