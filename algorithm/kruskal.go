package algorithm

import (
	"github.com/ymcvalu/ds_algo/data_structure/uf_set"
	"sort"
)

type Edge struct {
	From   int
	To     int
	Weight float64
}

func kruskal(n int, edges []Edge) ([]Edge, bool) {
	// n个顶点至少需要n-1条边
	if len(edges) < n-1 {
		return nil, false
	}
	ufs := uf_set.New(n)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	min := make([]Edge, 0, n-1)

	for i := 0; ufs.Count() > 1 && i < len(edges); i++ {
		e := edges[i]
		if !ufs.Connected(e.From, e.To) {
			ufs.Union(e.From, e.To)
			min = append(min, e)
		}

	}

	if ufs.Count() != 1 {
		return nil, false
	}

	return min, true
}
