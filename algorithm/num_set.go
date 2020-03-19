package algorithm

func findOneElem2(num []int) int {
	m := 0
	for _, n := range num {
		m ^= n
	}
	return m
}
