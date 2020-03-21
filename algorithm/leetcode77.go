package algorithm

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合

func combine(n int, k int) [][]int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	result := [][]int{}
	_combine(arr, k, nil, &result)
	return result
}

func _combine(arr []int, k int, path []int, result *[][]int) {
	if l := len(arr); l == 0 {
		return
	} else if l < k {
		return
	} else if l == k {
		p := path[:len(path):len(path)]
		*result = append(*result, append(p, arr...))
		return
	}

	_combine(arr[1:], k, path, result)
	path = append(path, arr[0])
	k--
	if k == 0 {
		*result = append(*result, append([]int(nil), path...))
	} else {
		_combine(arr[1:], k, path, result)
	}
}
