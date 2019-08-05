package avl_tree

type CompareFunc func(a, b interface{}) int

type Node struct {
	Key    interface{}
	Value  interface{}
	left   *Node
	right  *Node
	height int
}

func (n *Node) h() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) PreOrder(ns []*Node) []*Node {
	if n == nil {
		return ns
	}
	ns = append(ns, n)
	ns = n.left.PreOrder(ns)
	ns = n.right.PreOrder(ns)
	return ns
}

func (n *Node) rightRotate() *Node {
	left := n.left
	n.left = left.right
	left.right = n

	n.height = max(n.left.h(), n.right.h()) + 1
	left.height = max(left.left.h(), left.right.h()) + 1

	return left
}

func (n *Node) leftRotate() *Node {
	right := n.right
	n.right = right.left
	right.left = n

	n.height = max(n.left.h(), n.right.h()) + 1
	right.height = max(right.left.h(), right.right.h()) + 1

	return right
}

func (n *Node) reBalance(com CompareFunc) *Node {
	balance := n.left.h() - n.right.h()
	if balance > 1 {
		left := n.left
		if left.left.h() > left.right.h() {
			return n.rightRotate()
		} else {
			n.left = n.left.leftRotate()
			return n.rightRotate()
		}
	} else if balance < -1 {
		right := n.right
		if right.right.h() > right.left.h() {
			return n.leftRotate()
		} else {
			n.right = n.right.rightRotate()
			return n.leftRotate()
		}
	}

	n.height = max(n.left.h(), n.right.h()) + 1
	return n
}

func (n *Node) insert(nd *Node, com CompareFunc) *Node {
	if n == nil {
		return nd
	}
	ret := com(nd.Key, n.Key)
	if ret == 0 {
		return n
	}

	if ret < 0 {
		n.left = n.left.insert(nd, com)
	} else {
		n.right = n.right.insert(nd, com)
	}

	return n.reBalance(com)
}

func (n *Node) delete(key interface{}, com CompareFunc) *Node {
	if n == nil {
		return nil
	}
	ret := com(key, n.Key)
	if ret < 0 {
		n.left = n.left.delete(key, com)
	} else if ret > 0 {
		n.right = n.right.delete(key, com)
	} else {
		if n.left == nil || n.right == nil {
			if n.left == nil && n.right == nil {
				return nil
			} else if n.left != nil {
				return n.left
			} else if n.right != nil {
				return n.right
			}
		}

		del := n.left
		for del.right != nil {
			del = del.right
		}
		n.Key = del.Key
		n.Value = del.Value
		n.left = n.left.delete(del.Key, com)
	}

	// n.height = max(n.left.h(), n.right.h()) + 1

	return n.reBalance(com)
}

func (n *Node) find(key interface{}, com CompareFunc) *Node {
	if n == nil {
		return nil
	}
	ret := com(key, n.Key)
	if ret == 0 {
		return n
	}
	if ret < 0 {
		return n.left.find(key, com)
	} else {
		return n.right.find(key, com)
	}

}

type AVLTree struct {
	root *Node
	// =0: a == b
	// >0: a > b
	// <0：a < b
	compare CompareFunc
}

func New(compare CompareFunc) *AVLTree {
	return &AVLTree{
		compare: compare,
	}
}

func (t *AVLTree) Height() int {
	return t.root.h()
}

func (t *AVLTree) Insert(key, val interface{}) {
	nd := &Node{
		Key:    key,
		Value:  val,
		height: 1,
	}
	t.root = t.root.insert(nd, t.compare)
}

func (t *AVLTree) Find(key interface{}) interface{} {
	n := t.root.find(key, t.compare)
	if n == nil {
		return nil
	}
	return n.Value
}

func (t *AVLTree) Delete(key interface{}) {
	t.root = t.root.delete(key, t.compare)
}

func (t *AVLTree) PreOrder() (ns []*Node) {
	return t.root.PreOrder(ns)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
