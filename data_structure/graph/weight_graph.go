package graph

type Edge struct {
	To     int
	Weight float64
}

type WeightGraph struct {
	vn  int
	en  int
	dir bool
	adj [][]Edge
}

func NewWeightGraph(n int, dir bool) *WeightGraph {
	return &WeightGraph{
		vn:  n,
		dir: dir,
		adj: make([][]Edge, n),
	}
}

func (g *WeightGraph) NodeNum() int {
	return g.vn
}

func (g *WeightGraph) AddEdge(from, to int, w float64) {
	if from < 0 || to < 0 || from >= g.vn || to >= g.vn {
		panic("invalid params")
	}

	g.adj[from] = append(g.adj[from], Edge{
		To:     to,
		Weight: w,
	})

	if !g.dir {
		g.adj[to] = append(g.adj[to], Edge{
			To:     from,
			Weight: w,
		})
	}

	g.en++
}

func (g *WeightGraph) Adj(v int) []Edge {
	if v < 0 || v >= g.vn {
		panic("invalid param")
	}
	adj := g.adj[v]

	_adj := make([]Edge, len(adj))
	copy(_adj, adj)
	return _adj
}

func (g *WeightGraph) Reverse() *WeightGraph {
	if !g.dir {
		panic("g is not a directed graph")
	}
	r := &WeightGraph{
		vn:  g.vn,
		en:  g.en,
		adj: make([][]Edge, len(g.adj)),
		dir: g.dir,
	}

	for i := 0; i < len(g.adj); i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			r.adj[j] = append(r.adj[j], Edge{
				To:     i,
				Weight: r.adj[i][j].Weight,
			})
		}
	}

	return r
}
