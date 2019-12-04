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

	m := make([][]int, ln1+1)
	m[0] = make([]int, ln2+1)
	for i := 0; i < ln1+1; i++ {
		m[i] = make([]int, ln2+1)
	}
	for i := 1; i < ln1+1; i++ {
		m[i][0] = i
	}
	for i := 1; i < ln2+1; i++ {
		m[0][i] = i
	}

	for i := 1; i < ln1+1; i++ {

		for j := 1; j < ln2+1; j++ {
			if s1[i-1] == s2[j-1] {
				m[i][j] = m[i-1][j-1]
			} else {
				m[i][j] = min(min(m[i-1][j], m[i][j-1]), m[i-1][j-1]) + 1
			}
		}
	}

	return m[ln1][ln2]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
