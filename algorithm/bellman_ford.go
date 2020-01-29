package algorithm

import (
	"leetcode/data_structure/graph"
	"math"
)

// 单源最短路径，支持负权重边的DAG
// 只能用于没有负权重环的DAG
func BellmanFord(g *graph.WeightGraph, s int) []float64 {
	dist := make([]float64, g.NodeNum())
	for i := 0; i < g.NodeNum(); i++ {
		if i != s {
			dist[i] = math.MaxFloat64
		}
	}
	// 记录最短路径，可以使用该数组进行负权重环检测
	edgeTo := make([]int, g.NodeNum())
	l := NewLinkedList()
	set := map[int]bool{}

	l.Push(s)
	set[s] = true
	for !l.IsEmpty() {
		v := l.Pop().(int)
		delete(set, v)

		for _, adj := range g.Adj(v) {
			if dist[adj.To] > dist[v]+adj.Weight {
				dist[adj.To] = dist[v] + adj.Weight
				edgeTo[adj.To] = v
				if !set[adj.To] {
					l.Push(adj.To)
					set[adj.To] = true
				}
			}
		}

	}

	return dist
}
