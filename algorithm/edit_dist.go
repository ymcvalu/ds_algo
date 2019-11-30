package algorithm

// the edit-dist from s1 to s2
func editDist(s1, s2 string) int {
	ln1 := len(s1)
	ln2 := len(s2)
	if ln1 == 0 {
		return ln2
	} else if ln2 == 0 {
		return ln1
	}

	m := make([][]int, ln1)
	m[0] = make([]int, ln2)

	for i := 0; i < ln1; i++ {
		m[i] = make([]int, ln2)

		for j := 0; j < ln2; j++ {
			if s1[i] == s2[j] {
				if i == 0 && j == 0 {
					m[i][j] = 0
				} else if i == 0 {
					m[i][j] = m[i][j-1]
				} else if j == 0 {
					m[i][j] = m[i-1][j]
				} else {
					m[i][j] = m[i-1][j-1]
				}
			} else {
				if i == 0 && j == 0 {
					m[i][j] = 1
				} else if i == 0 {
					m[i][j] = m[i][j-1] + 1
				} else if j == 0 {
					m[i][j] = m[i-1][j] + 1
				} else {
					m[i][j] = min(min(m[i-1][j], m[i][j-1]), m[i-1][j-1]) + 1
				}
			}
		}
	}

	return m[ln1-1][ln2-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
