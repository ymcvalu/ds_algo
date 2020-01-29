package segment_tree

import (
	"testing"
)

func TestSegmentTreeSum(t *testing.T) {
	st := New([]interface{}{1, 2, 3, 4, 5, 6}, func(i, j interface{}) interface{} {
		return i.(int) + j.(int)
	})

	t.Log(st.Reduce(0, 5))
	t.Log(st.Reduce(1, 2))
	t.Log(st.Reduce(2, 5))
	t.Log(st.Reduce(4, 5))
	t.Log(st.Reduce(2, 4))
	t.Log(st.Reduce(3, 5))
	t.Log(st.Reduce(1, 3))
	t.Log(st.Reduce(0, 3))
}

func TestSegmentTreeMul(t *testing.T) {
	st := New([]interface{}{1, 2, 3, 4, 5, 6}, func(i, j interface{}) interface{} {
		return i.(int) * j.(int)
	})

	t.Log(st.Reduce(0, 5))
	t.Log(st.Reduce(1, 2))
	t.Log(st.Reduce(2, 5))
	t.Log(st.Reduce(4, 5))
	t.Log(st.Reduce(2, 4))
	t.Log(st.Reduce(3, 5))
	t.Log(st.Reduce(1, 3))
	t.Log(st.Reduce(0, 3))
}

func TestSegmentTreeMax(t *testing.T) {
	st := New([]interface{}{1, 2, 3, 4, 5, 6}, func(i, j interface{}) interface{} {
		if i.(int) >= j.(int) {
			return i
		}
		return j
	})
	t.Log(st.t)
	t.Log(st.Reduce(0, 5))
	t.Log(st.Reduce(1, 2))
	t.Log(st.Reduce(2, 5))
	t.Log(st.Reduce(4, 5))
	t.Log(st.Reduce(2, 4))
	t.Log(st.Reduce(3, 5))
	t.Log(st.Reduce(1, 3))
	t.Log(st.Reduce(0, 3))
}

func TestSegmentTreeMin(t *testing.T) {
	st := New([]interface{}{1, 2, 3, 4, 5, 6}, func(i, j interface{}) interface{} {
		if i.(int) <= j.(int) {
			return i
		}
		return j
	})

	t.Log(st.Reduce(0, 5))
	t.Log(st.Reduce(1, 2))
	t.Log(st.Reduce(2, 5))
	t.Log(st.Reduce(4, 5))
	t.Log(st.Reduce(2, 4))
	t.Log(st.Reduce(3, 5))
	t.Log(st.Reduce(1, 3))
	t.Log(st.Reduce(0, 3))
}

func TestSegmentTreeUpdate(t *testing.T) {
	st := New([]interface{}{1, 2, 3, 4, 5, 6}, func(i, j interface{}) interface{} {
		return i.(int) + j.(int)
	})
	t.Log(st.t)
	st.Update(0, 7)
	t.Log(st.t)
	st.Update(3, 8)
	t.Log(st.t)
	st.Update(5, 1)
	t.Log(st.t)
	st.Update(0, 1)
	st.Update(3, 4)
	st.Update(5, 6)
	t.Log(st.t)
}
