package algorithm

/**
输入两个整数 n 和 m，从数列1，2，3.......n 中 随意取几个数,使其和等于 m ,
要求将其中所有的可能组合列出来.
*/

func SumM(n, m int) [][]int {
	ret := make([][]int, 0)
	sumMBacktrace(1, n, m, 0, nil, &ret)
	return ret
}

func sumMBacktrace(i, n, m, s int, path []int, ret *[][]int) {
	if i > n {
		return
	}

	if s+i < m {
		// 使用当前i
		sumMBacktrace(i+1, n, m, s+i, append(path, i), ret)
	} else if s+i == m {
		r := make([]int, len(path)+1)
		copy(r, append(path, i))
		*ret = append(*ret, r)
	}
	sumMBacktrace(i+1, n, m, s, path, ret) // 不使用i
}
