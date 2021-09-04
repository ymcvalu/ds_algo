package algorithm

import "github.com/ymcvalu/ds_algo/data_structure/graph"

func buildGraph(dir bool, n int, pairs ...int) *graph.Graph {
	g := graph.New(n, dir)
	for i := 0; i < len(pairs); i += 2 {
		g.AddEdge(pairs[i], pairs[i+1])
	}
	return g
}

func buildWeightGraph(dir bool, n int, pairs []Edge) *graph.WeightGraph {
	g := graph.NewWeightGraph(n, dir)
	for _, p := range pairs {
		g.AddEdge(p.From, p.To, p.Weight)
	}
	return g
}
