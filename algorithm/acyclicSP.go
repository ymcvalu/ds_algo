package algorithm

import "github.com/ymcvalu/ds_algo/data_structure/graph"

// 对于无环的加权有向图，只要按照拓扑顺序处理无环有向图的顶点即可
// 既可以处理最短路径，也可以处理最长路径
// 关键路径：该路径上的最早开始时间于最晚开始时间一致，实际上就是最长路径
// AOE网关键路径：先计算从起点到终点的最长路径，同时也计算最早开始时间，然后从终点向起点回溯，计算最晚开始时间
func AcyclicSP(g *graph.WeightGraph) []float64 {
	// TODO
	return nil
}
