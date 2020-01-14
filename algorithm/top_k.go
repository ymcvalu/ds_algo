package algorithm

func topk(arr []int, k int) int {
	if k < 1 || len(arr) < k {
		return -1
	}

	p := _partition(arr)
	if p+1 == k {
		return arr[p]
	}

	if p+1 < k {
		return topk(arr[p+1:], k-p-1)
	}

	return topk(arr[:p], k)
}

func _partition(arr []int) int {
	l := len(arr)
	if l == 1 {
		return 0
	}

	var i = 0
	var p = i + 1
	var q = l - 1

	for {
		for p <= q && arr[p] > arr[i] {
			p++
		}
		for q >= p && arr[q] < arr[i] {
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
