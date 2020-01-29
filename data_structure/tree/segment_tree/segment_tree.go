package segment_tree

import "math"

type Tree struct {
	n int
	t []interface{}
	r Reduce
}

type Reduce func(i, j interface{}) interface{}

func New(arr []interface{}, r Reduce) *Tree {
	n := len(arr)
	if n == 0 || r == nil {
		panic("invalid params")
	}

	h := math.Ceil(math.Log2(float64(n)))

	t := &Tree{
		n: n,
		t: make([]interface{}, int(math.Pow(2, h+1)-1)),
		r: r,
	}

	t.build(arr, 0, 0, n-1)
	return t
}

func (t *Tree) build(arr []interface{}, si int, l, r int) {
	if l == r {
		t.t[si] = arr[l]
		return
	}

	mid := (l + r) / 2
	t.build(arr, si*2+1, l, mid)
	t.build(arr, si*2+2, mid+1, r)
	t.t[si] = t.r(t.t[si*2+1], t.t[si*2+2])
}

func (t *Tree) Reduce(l, r int) interface{} {
	if l < 0 || r < 0 || l >= t.n || r >= t.n || l > r {
		return nil
	}

	return t.reduce(0, 0, t.n-1, l, r)

	return nil
}

func (t *Tree) reduce(si, rl, rr, l, r int) interface{} {
	if rl == l && rr == r {
		return t.t[si]
	}

	mid := (rl + rr) / 2

	if l > mid {
		return t.reduce(si*2+2, mid+1, rr, l, r)
	} else if r <= mid {
		return t.reduce(si*2+1, rl, mid, l, r)
	} else {
		return t.r(t.reduce(si*2+1, rl, mid, l, mid), t.reduce(si*2+2, mid+1, rr, mid+1, r))
	}
}

func (t *Tree) Update(idx int, v interface{}) {
	if idx < 0 || idx >= t.n {
		panic("invalid param")
	}
	t.update(0, 0, t.n-1, idx, v)
}

func (t *Tree) update(si int, l, r int, idx int, v interface{}) {
	if l == r {
		t.t[si] = v
		return
	}

	mid := (l + r) / 2
	if idx > mid {
		t.update(si*2+2, mid+1, r, idx, v)
	} else {
		t.update(si*2+1, l, mid, idx, v)
	}
	t.t[si] = t.r(t.t[si*2+1], t.t[si*2+2])
}
