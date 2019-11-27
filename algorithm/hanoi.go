package algorithm

import "fmt"

func Hanoi(n int) {
	hanoi("a", "b", "c", n)
}

func hanoi(from, to, by string, n int) {
	if n == 1 {
		fmt.Printf("move %d from %s to %s\n", n, from, to)
		return
	}
	hanoi(from, by, to, n-1)
	fmt.Printf("move %d from %s to %s\n", n, from, to)
	hanoi(by, to, from,n-1)
}
