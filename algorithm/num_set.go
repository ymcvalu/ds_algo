package algorithm

func findOneElem3(num []int) int {
	m := 0
	for _, n := range num {
		m = ^(m ^ n)
	}
	return m
}

func findOneElem2(num []int) int {
	m := 0
	for _, n := range num {
		m ^= n
	}
	return m
}
