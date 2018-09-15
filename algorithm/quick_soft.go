package algorithm

func QuickSoft(arr []int) {
	l := len(arr)
	if l == 1 {
		return
	}
	mid := partition(arr)
	if mid > 0 {
		QuickSoft(arr[:mid])
	}
	if mid < l-1 {
		QuickSoft(arr[mid+1:])
	}
}

func partition(arr []int) int {
	l := len(arr)
	if l == 1 {
		return 0
	}

	var i = 0
	var p = i + 1
	var q = l - 1

	for {
		for p <= q && arr[p] < arr[i] {
			p++
		}
		for q >= p && arr[q] > arr[i] {
			q--
		}
		if p >= q {
			break
		}
		arr[p], arr[q] = arr[q], arr[p]
		p++
		q--

	}

	arr[i], arr[q] = arr[q], arr[i]

	return q
}
