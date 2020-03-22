package algorithm

import (
	"math/rand"
	"sort"
	"time"
)

func randWeight(w []int) int {
	rand.Seed(time.Now().UnixNano())
	sum := 0
	for i := 0; i < len(w); i++ {
		sum += w[i]
	}
	r := rand.Intn(sum)
	sort.Ints(w)
	c := 0
	for i := 0; i < len(w); i++ {
		c += w[i]
		if r < c {
			return w[i]
		}
	}
	return -1
}
