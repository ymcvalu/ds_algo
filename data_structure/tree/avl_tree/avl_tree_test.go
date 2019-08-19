package avl_tree

import "testing"

func TestInsert(t *testing.T) {
	tree := New(func(a, b interface{}) int {
		ai := a.(int)
		bi := b.(int)
		if ai > bi {
			return 1
		} else if ai < bi {
			return -1
		}
		return 0
	})
	tree.Insert(6, "")
	tree.Insert(8, "")
	tree.Insert(3, "")
	tree.Insert(9, "9")
	tree.Insert(7, "")
	tree.Insert(1, "")
	tree.Insert(2, "")
	tree.Insert(4, "4")
	tree.Insert(5, "")
	tree.Insert(16, "")
	tree.Insert(13, "")
	tree.Insert(20, "")

	t.Log(tree.Find(9))
	t.Log(tree.Find(4))

	tree.Delete(4)
	tree.Delete(0)
	tree.Delete(20)
	tree.Delete(6)
	tree.Delete(7)
	tree.Delete(1)
	tree.Delete(5)
	tree.Delete(4)

	nodes := tree.Slice()

	for _, n := range nodes {
		t.Log(n.height, n.Key)
	}

	tree.Walk(func(k, v interface{}) bool {
		t.Log(k)
		if k==13{
			return false
		}
		return true
	})

}
