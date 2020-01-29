package algorithm

import (
	"leetcode/data_structure/graph"
	"math"
	"unsafe"
)

// dijkstra单源最短路径算法只能用于求解边的权重非负的情况
//       2
//   a ----- c
//    \     /
//   1 \   / -2
//       b
// 如上，起点为a，a->b 小于 a->c，因此根据算法，确定a->b的最短距离为1，但是实际上，a->b的最短距离为a->c->b，距离为0
// dijkstra假设添加一条边会使路径变的更长，但是如果添加的是负权重只会使路径更短
// dijkstra本质上是贪心算法

var MAX_INT = (1<<unsafe.Sizeof(int(0))*8 - 1) - 1

func dijkstraSP(g *graph.WeightGraph, s int) []float64 {
	dist := make([]float64, g.NodeNum())
	edgeTo := make([]int, g.NodeNum())
	for i := 0; i < len(dist); i++ {
		if i != s {
			dist[i] = math.MaxFloat64
		}
	}

	type Node struct {
		V    int
		Dist float64
	}

	pq := NewPriorityQueue()

	for _, adj := range g.Adj(s) {
		dist[adj.To] = adj.Weight
		edgeTo[adj.To] = adj.To
		pq.Push(adj.To, adj.Weight)
	}

	for {
		v, has := pq.Pop()
		if !has {
			break
		}

		for _, adj := range g.Adj(v) {
			if dist[adj.To] > dist[v]+adj.Weight {
				dist[adj.To] = dist[v] + adj.Weight
				edgeTo[adj.To] = v
				pq.Push(adj.To, dist[adj.To])
			}
		}
	}

	return dist
}
