package algorithm

func gcd(x, y int) int {
	for {
		if x == y {
			return x
		}

		if x < y {
			x, y = y, x
		}
		x = x - y
		if x == 1 {
			return 1
		}
	}
}
