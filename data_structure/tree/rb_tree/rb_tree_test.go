package rb_tree

import (
	"math/rand"
	"testing"
	"time"
)

func TestRbTree(t *testing.T) {
	N := 1000000
	D := 666666

	set := make(map[int]struct{}, N)
	rb := New(func(i, j interface{}) int {
		return i.(int) - j.(int)
	})

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < N; i++ {
		n := r.Intn(N << 1)
		set[n] = struct{}{}
		rb.Ins(n, struct{}{})
	}

	if !rb.IsRbTree() {
		t.Errorf("not a rb tree after ins")
		return
	}

	for k := range set {
		if _, has := rb.Search(k); !has {
			t.Errorf("less key %d after ins", k)
			return
		}
	}

	d := D
	for k := range set {
		if d == 0 {
			break
		}
		delete(set, k)
		if rb.Del(k) != struct{}{} {
			t.Errorf("less key %d when del", k)
			rb.PreOder()
			return
		}
		d--
	}

	for k := range set {
		if _, has := rb.Search(k); !has {
			t.Errorf("less key %d after del", k)
			break
		}
	}

	t.Log(rb.IsRbTree())
}

func TestDel(t *testing.T) {
	var set = []int{6, 3, 8, 1, 2, 7, 9, 4, 5, 10}
	rb := New(func(i, j interface{}) int {
		return i.(int) - j.(int)
	})

	for _, s := range set {
		rb.Ins(s, struct{}{})
	}

	rb.Del(3)
	rb.Del(2)
	rb.Del(1)
	rb.Del(10)
	rb.Del(6)
	rb.Del(10)

	rb.PreOder()

	t.Log(rb.IsRbTree())

}
