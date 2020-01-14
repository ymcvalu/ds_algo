package algorithm

func quickSort3(arr []int) {
	ln := len(arr)
	if ln <= 2 {
		return
	}
	l, r := partition3(arr)

	if len(arr[:l]) > 2 {
		quickSort3(arr[:l])
	}
	if len(arr[r:]) > 2 {
		quickSort3(arr[r:])
	}

}

func partition3(arr []int) (l, r int) {

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

		if arr[q] == v {
			arr[p], arr[q] = arr[q], arr[p]
		} else {
			arr[l], arr[q] = arr[q], arr[l]
			l++
		}
	}

	r = p

	return
}
