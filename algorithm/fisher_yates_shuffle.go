package algorithm

import (
	"math/rand"
	"time"
)

// 等概率洗牌算法
func FisherYatesShuffle(n int) []int {
	num := make([]int, n)
	for i := 1; i < n; i++ {
		num[i] = i
	}
	rand.Seed(time.Now().UnixNano())

	// 首先选举第n个数的时候，概率为1/n，剩余n-1个数没有被选中的概率为(n-1)/n
	// 接下来选举第n-1个数的时候，概率为 ((n-1)/n)*(1/(n-1))= 1/n
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		num[i], num[j] = num[j], num[i]
	}
	return num
}
