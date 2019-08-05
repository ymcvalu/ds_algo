package algorithm

func LCS(s1, s2 string) string {
	ln1, ln2 := len(s1), len(s2)
	max := ln1
	if max < ln2 {
		max = ln2
	}

	ln1++
	ln2++

	// 构造二维数组，空间换时间
	m := make([][]int, ln1)
	for i := 0; i < ln1; i++ {
		m[i] = make([]int, ln2)
	}

	for i := 1; i < ln1; i++ {
		for j := 1; j < ln2; j++ {
			if s1[i-1] == s2[j-1] {
				m[i][j] = m[i-1][j-1] + 1
			} else {
				v := m[i-1][j-1]
				if v < m[i][j-1] {
					v = m[i][j-1]
				}
				if v < m[i-1][j] {
					v = m[i-1][j]
				}
				m[i][j] = v
			}
		}
	}

	// 回溯
	i, j := ln1-1, ln2-1
	len := m[i][j]
	ret := make([]byte, len)

	for len > 0 {
		if n := m[i][j]; n > m[i-1][j-1] {
			len--
			ret[len] = s1[i-1]
			i--
			j--
		} else if n > m[i-1][j] {
			len--
			ret[len] = s1[i-1]
			i--
		} else if n > m[i][j-1] {
			len--
			ret[len] = s1[i-1]
			j--
		} else {
			i--
			j--
		}
	}

	return string(ret)
}
