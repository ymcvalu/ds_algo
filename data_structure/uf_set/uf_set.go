package uf_set

// 并查集
type UFSet struct {
	n  int
	id []int
	c  int
}

func New(n int) *UFSet {
	s := &UFSet{
		n:  n,
		id: make([]int, n),
		c:  n,
	}
	for i := 0; i < n; i++ {
		s.id[i] = i
	}
	return s
}

func (s *UFSet) Find(p int) int {
	if p < 0 || p >= s.n {
		panic("out of set")
	}

	// 路径压缩
	g := s.id[p]
	if g != p {
		g = s.Find(g)
		s.id[p] = g
	}
	return g
}

func (s *UFSet) Union(p, q int) {
	if p < 0 || p >= s.n || q < 0 || q >= s.n {
		panic("out of set")
	}

	gp := s.Find(p)
	gq := s.Find(q)

	if gp != gq {
		s.id[gq] = gp
		s.c--
	}
}

func (s *UFSet) Connected(p, q int) bool {
	if p < 0 || p >= s.n || q < 0 || q >= s.n {
		panic("out of set")
	}
	return s.Find(p) == s.Find(q)
}

func (s *UFSet) Count() int {
	return s.c
}
