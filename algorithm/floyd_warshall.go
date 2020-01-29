package algorithm

import (
	"leetcode/data_structure/graph"
	"math"
)

// 无负权环的DAG，任意两点之间的最短距离
// 如果需要检测负权环，可用添加一个追踪上一节点的Matrix，然后通过该Matrix进行检测
func FloydWarshall(g *graph.WeightGraph) [][]float64 {
	dp := make([][]float64, g.NodeNum())

	for i := 0; i < g.NodeNum(); i++ {
		dp[i] = make([]float64, g.NodeNum())

		for j := 0; j < g.NodeNum(); j++ {
			if i != j {
				dp[i][j] = math.MaxFloat64
			}
		}
	}

	for v := 0; v < g.NodeNum(); v++ {
		for _, adj := range g.Adj(v) {
			dp[v][adj.To] = adj.Weight
		}
	}

	for k := 0; k < g.NodeNum(); k++ {
		for i := 0; i < g.NodeNum(); i++ {
			for j := 0; j < g.NodeNum(); j++ {
				if dist := dp[i][k] + dp[k][j]; dp[i][j] > dist {
					dp[i][j] = dist
				}
			}
		}
	}

	return dp
}
