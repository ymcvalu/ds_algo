package algorithm

func genBC(pattern string) []int {
	bc := make([]int, 256)
	for i := range bc {
		bc[i] = -1
	}

	for i, c := range pattern {
		bc[int(c)] = i
	}
	return bc
}

func genGS(pattern string) ([]int, []bool) {
	prefix := make([]bool, len(pattern))
	suffix := make([]int, len(pattern))
	for i := range suffix {
		suffix[i] = -1
	}

	lp := len(pattern)

	for i := 0; i < lp-1; i++ {
		j, k := i, 0

		for j >= 0 && pattern[j] == pattern[lp-1-k] {
			k++
			suffix[k] = j
			j--
		}
		if j == -1 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}

func moveByGS(j, m int, suffix []int, prefix []bool) int {
	k := m - j - 1       // 已经匹配的后缀字符串为[j+1,m-1]，k为字符个数
	if suffix[k] != -1 { // 字符串中包含该后缀
		return j - suffix[k] + 1
	}

	// 查找后缀的字串是否是模式串的前缀
	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] == true {
			return r
		}
	}
	return m
}

func bm(str, pattern string) int {
	bc := genBC(pattern)
	suffix, prefix := genGS(pattern)

	i, j := 0, 0
	for ls, lp := len(str), len(pattern); i <= ls-lp; {
		for j = lp - 1; j >= 0 && str[i+j] == pattern[j]; j-- {
		}
		if j < 0 {
			return i // match
		}

		x := j - bc[int(str[i+j])]
		y := 0
		if j < lp-1 {
			y = moveByGS(j, lp, suffix, prefix)
		}
		i += max(x, y)
	}

	return -1
}
