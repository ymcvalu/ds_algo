package algorithm

func quickSort3a(arr []int) {
	ln := len(arr)
	if ln < 2 {
		return
	}
	l, r := partition3a(arr)
	if len(arr[:l]) > 1 {
		quickSort3a(arr[:l])
	}
	if len(arr[r:]) > 1 {
		quickSort3a(arr[r:])
	}

}

func partition3a(arr []int) (l, r int) {

	ln := len(arr)
	if ln < 2 {
		return 0, ln
	}

	v := arr[0]
	p := 1
	q := ln - 1

	for {
		for p <= q && arr[p] <= v {
			if arr[p] == v {
				p++
				continue
			}

			arr[l], arr[p] = arr[p], arr[l]

			l++
			p++
		}

		for p <= q && arr[q] > v {
			q--
		}

		if p > q {
			break
		}

		arr[p], arr[q] = arr[q], arr[p]
		if arr[p] < v {
			arr[l], arr[p] = arr[p], arr[l]
			l++
		}

		p++
		q--

	}

	r = p

	return
}

func quickSort3b(arr []int) {
	ln := len(arr)
	if ln < 2 {
		return
	}
	l, r := partition3b(arr)
	if len(arr[:l]) > 1 {
		quickSort3b(arr[:l])
	}
	if len(arr[r:]) > 1 {
		quickSort3b(arr[r:])
	}

}
func partition3b(arr []int) (l, r int) {
	ln := len(arr)
	l = 0
	r = ln - 1
	v := arr[0]

	for i := 1; i <= r; {
		if arr[i] < v {
			arr[l], arr[i] = arr[i], arr[l]
			l++
			i++
		} else if arr[i] > v {
			arr[r], arr[i] = arr[i], arr[r]
			r--
		} else {
			i++
		}
	}

	r += 1
	return
}
