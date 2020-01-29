package uf_set

import "testing"

func TestUFSet(t *testing.T) {
	set := New(10)
	t.Log(set.Connected(1, 3), false)
	set.Union(1, 3)
	t.Log(set.Connected(1, 3), true)
	t.Log(set.Connected(1, 4), false)
	set.Union(1, 4)
	t.Log(set.Connected(3, 4), true)
	t.Log(set.Count(), 8)
	t.Log(set.Find(1), set.Find(3), set.Find(4))
	t.Log(set.Find(9))
	set.Union(8, 9)
	set.Union(0, 1)
	t.Log(set.Find(8), set.Find(9))
}
