package algorithm

import (
	"fmt"
	"github.com/ymcvalu/ds_algo/data_structure/graph"
	"github.com/ymcvalu/ds_algo/data_structure/stack"
)

func topologicalSortA(g *graph.Graph) ([]int, bool) {
	degree := make([]int, g.NodeNum())
	for i := 0; i < g.NodeNum(); i++ {
		adj := g.Adj(i)
		for _, v := range adj {
			degree[v]++
		}
	}
	ns := make([]int, 0, g.NodeNum())

	var vs []int
	// bfs;可以使用heap优化
	for i := 0; i < len(degree); {
		for j := 0; j < len(degree); j++ {
			if degree[j] == 0 {
				vs = append(vs, j)
				degree[j] = -1
				i++
			}
		}
		if len(vs) == 0 {
			return nil, false
		}
		for _, v := range vs {
			ns = append(ns, v)
			adj := g.Adj(v)
			for _, a := range adj {
				degree[a]--
			}
		}

		vs = vs[:0]
	}

	return ns, true
}

// 只适用于DAG，如果存在环应该无法进行拓扑排序才对
func topologicalSortB(g *graph.Graph) ([]int, bool) {
	degree := make([]int, g.NodeNum())
	for i := 0; i < g.NodeNum(); i++ {
		adj := g.Adj(i)
		for _, v := range adj {
			degree[v]++
		}
	}

	ns := make([]int, 0, g.NodeNum())
	for i, v := range degree {
		if v == 0 {
			ns = append(ns, i)
		}
	}

	if len(ns) == 0 {
		fmt.Println("1")
		return nil, false
	}

	marked := make([]bool, g.NodeNum())
	stk := stack.New()
	var dfs func(int)
	dfs = func(v int) {
		marked[v] = true
		for _, n := range g.Adj(v) {
			if !marked[n] {
				dfs(n)
			}
		}
		stk.Push(v)
	}

	for _, v := range ns {
		dfs(v)
	}

	nodes := make([]int, 0, g.NodeNum())
	for {
		v := stk.Pop()
		if v == nil {
			break
		}
		nodes = append(nodes, v.(int))
	}
	return nodes, true
}
