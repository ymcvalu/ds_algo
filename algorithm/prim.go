package algorithm

import (
	"fmt"
	"math"
)

func Prim(graph [][]int64) {
	if len(graph) <= 1 {
		return
	}

	dist := make([]int64, len(graph))
	parent := make([]int, len(graph))
	for i := 1; i < len(graph); i++ {
		dist[i] = graph[0][i]
	}

	// 可以使用heap优化
	for {
		idx := findMin(dist)
		if idx == -1 {
			break
		}
		fmt.Printf("chose edge from node %d to node %d, dist is %d\n", parent[idx], idx, dist[idx])
		dist[idx] = 0

		for i := 0; i < len(dist); i++ {
			if dist[i] > 0 && graph[idx][i] < dist[i] {
				dist[i] = graph[idx][i]
				parent[i] = idx
			}
		}

	}

}

func findMin(dist []int64) int {
	min := int64(math.MaxInt64)
	idx := -1
	for i := range dist {
		if dist[i] > 0 && dist[i] < min {
			min = dist[i]
			idx = i
		}
	}
	return idx
}
