package algorithm

func FibonacciA(n int) int {
	v1 := 0
	v2 := 1
	for i := 2; i <= n; i++ {
		v2, v1 = v2+v1, v2
	}
	return v2
}

func FibonacciB(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibonacciB(n-1) + FibonacciB(n-2)
}

func FibonacciC(n int) int {
	var fn func(int, []int) int
	fn = func(n int, t []int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		if t[n] > 0 {
			return t[n]
		}
		t[n] = fn(n-1, t) + fn(n-2, t)
		return t[n]
	}
	return fn(n, make([]int, n+1))
}
