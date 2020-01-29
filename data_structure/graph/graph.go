package graph

type Graph struct {
	vn  int
	en  int
	adj [][]int
	dir bool
}

func New(n int, dir bool) *Graph {
	return &Graph{
		vn:  n,
		en:  0,
		adj: make([][]int, n),
		dir: dir,
	}
}

func (g *Graph) NodeNum() int {
	return g.vn
}

func (g *Graph) AddEdge(from, to int) {
	if from < 0 || to < 0 || from >= g.vn || to >= g.vn {
		panic("invalid params")
	}
	g.adj[from] = append(g.adj[from], to)
	if !g.dir {
		g.adj[to] = append(g.adj[to], from)
	}
	g.en++
}

func (g *Graph) Adj(v int) []int {
	if v < 0 || v >= g.vn {
		panic("invalid param")
	}
	adj := g.adj[v]

	_adj := make([]int, len(adj))
	copy(_adj, adj)
	return _adj
}

func (g *Graph) Reverse() *Graph {
	if !g.dir {
		panic("g is not a directed graph")
	}
	r := &Graph{
		vn:  g.vn,
		en:  g.en,
		adj: make([][]int, len(g.adj)),
		dir: g.dir,
	}

	for i := 0; i < len(g.adj); i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			r.adj[j] = append(r.adj[j], i)
		}
	}

	return r
}
