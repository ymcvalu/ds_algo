package algorithm

/**
Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

For example,
[2,3,4], the median is 3

[2,3], the median is (2 + 3) / 2 = 2.5

Design a data structure that supports the following two operations:
  - void addNum(int num) - Add a integer number from the data stream to the data structure.
  - double findMedian() - Return the median of all elements so far.

Example:
  addNum(1)
  addNum(2)
  findMedian() -> 1.5
  addNum(3)
  findMedian() -> 2
*/

type MedianFinder struct {
	node *AVLTree
	cnt  int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		node: NewAVLTree(func(a, b int) int {
			if a > b {
				return 1
			} else if a == b {
				return 0
			} else {
				return -1
			}
		}),
	}
}

func (this *MedianFinder) AddNum(num int) {
	val := this.node.Find(num)
	if val == 0 {
		this.node.Insert(num, 1)
	} else {
		this.node.Insert(num, val+1)
	}
	this.cnt++
}

func (this *MedianFinder) FindMedian() float64 {
	var (
		even = this.cnt%2 == 0
		i1   = 0
		i2   = 0
		n    = 0
	)

	this.node.Walk(func(k, v int) bool {
		if even {
			if n == this.cnt/2 {
				i2 = k
				return false
			}
			n += v
			if n == this.cnt/2 {
				i1 = k
				return true
			} else if n > this.cnt/2 {
				i1 = k
				i2 = i1
				return false
			}
		} else {
			n += v
			if n > this.cnt/2 {
				i1 = k
				i2 = i1
				return false
			}
		}
		return true
	})

	return float64(i1+i2) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// val tree

type CompareFunc func(a, b int) int

type Node struct {
	Key    int
	Value  int
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

func (n *Node) Walk(f func(k, v int) bool) bool {
	if n == nil {
		return true
	}
	if !n.left.Walk(f) {
		return false
	}
	if !f(n.Key, n.Value) {
		return false
	}
	if !n.right.Walk(f) {
		return false
	}
	return true
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
		n.Value = nd.Value
		return n
	}

	if ret < 0 {
		n.left = n.left.insert(nd, com)
	} else {
		n.right = n.right.insert(nd, com)
	}

	return n.reBalance(com)
}

func (n *Node) find(key int, com CompareFunc) *Node {
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

func NewAVLTree(compare CompareFunc) *AVLTree {
	return &AVLTree{
		compare: compare,
	}
}

func (t *AVLTree) Insert(key, val int) {
	nd := &Node{
		Key:    key,
		Value:  val,
		height: 1,
	}
	t.root = t.root.insert(nd, t.compare)
}

func (t *AVLTree) Find(key int) int {
	n := t.root.find(key, t.compare)
	if n == nil {
		return 0
	}
	return n.Value
}

func (t *AVLTree) Walk(f func(k, v int) bool) {
	t.root.Walk(f)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
