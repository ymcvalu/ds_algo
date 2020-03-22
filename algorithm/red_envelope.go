package algorithm

import (
	"math/rand"
	"sort"
)

// 红包分割算法

// 二倍均值法：每次抢到的金额 = rand(0, M / N X 2)
func divideRedEnvelope2xMean(m int, n int) []int {
	pkgs := make([]int, 0, n)
	for i := n; i > 0; i-- {
		if i == 1 {
			pkgs = append(pkgs, m)
		} else {
			c := 0
			for c == 0 {
				c = rand.Intn(2 * m / i)
			}

			m -= c
			pkgs = append(pkgs, c)
		}
	}
	return pkgs
}

// 线性分割法
func divideRedEnvelopeLinearSplit(m int, n int) []int {
	pm := map[int]bool{}
	// 分割点
	for i := 0; i < n-1; {
		p := rand.Intn(m)
		if p == 0 || pm[p] {
			continue
		}
		pm[p] = true
		i++
	}

	ps := make([]int, 0, len(pm))
	for p := range pm {
		ps = append(ps, p)
	}

	sort.Ints(ps)

	pkgs := make([]int, 0, n)

	pkgs = append(pkgs, ps[0])
	for i := 1; i < len(ps); i++ {
		pkgs = append(pkgs, ps[i]-ps[i-1])
	}

	pkgs = append(pkgs, m-ps[len(ps)-1])

	return pkgs
}
