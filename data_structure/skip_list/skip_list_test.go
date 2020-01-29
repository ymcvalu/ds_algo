package skip_list

import "testing"

func print(t *testing.T, l *SkipList) {
	x := &l.head
	for x != nil {
		h := len(x.next)
		next := make([]interface{}, h)
		for i := 0; i < h; i++ {
			if x.next[i] != nil {
				next[i] = x.next[i].key
			}
		}
		t.Log(x.key, h, next)
		x = x.next[0]
	}
}

func TestSkipList(t *testing.T) {
	l := New(5, func(i, j interface{}) int {
		return i.(int) - j.(int)
	})

	l.Insert(5, nil)
	l.Insert(3, nil)
	l.Insert(7, nil)
	l.Insert(4, nil)
	l.Insert(1, nil)
	l.Insert(2, nil)
	l.Insert(8, nil)

	t.Log(l.Get(3))
	t.Log(l.Get(1))
	t.Log(l.Get(9))
	t.Log(l.Get(8))
	t.Log(l.Get(0))

	print(t, l)

	l.Delete(1)
	l.Delete(5)
	l.Delete(8)
	l.Delete(3)
	t.Log(l.Get(0))
	t.Log(l.Get(1))
	t.Log(l.Get(4))
	t.Log(l.Get(3))
	t.Log(l.Get(5))
	t.Log(l.Get(8))
	t.Log(l.Get(7))

	print(t, l)
}
