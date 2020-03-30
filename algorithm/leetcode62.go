package algorithm

/**
0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。

例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
*/

func lastRemainingRecursive(n int, m int) int {
	var f func(n, m int) int
	f = func(n, m int) int {
		if n == 1 {
			return 0
		}
		x := f(n-1, m)
		return (m + x) % n
	}
	return f(n, m)
}

func lastRemaining(n, m int) int {
	f := 0
	for i := 2; i != n+1; i++ {
		f = (m + f) % i
	}
	return f
}
