package rb_tree

import (
	"fmt"
)

// >0 when i>j
// =0 when i=j
// <0 when i<j
type Compare func(i, j interface{}) int

type color uint8

const (
	RED color = iota
	BLACK
)

type node struct {
	k     interface{}
	v     interface{}
	p     *node
	l     *node
	r     *node
	color color
}

// the leaves of rb-tree is always black and Nil
var Nil = &node{
	color: BLACK,
}

func init() {
	Nil.p = Nil
	Nil.l = Nil
	Nil.r = Nil
}

func (n *node) afterIns() *node {
	c := n
	for c.p != Nil && c.color == RED && c.p.color == RED {
		if u := c.uncle(); u.color == RED {
			// split the full 4-node
			c.p.color = BLACK
			u.color = BLACK
			c.p.p.color = RED
			c = c.p.p
		} else {
			lb := c.p == c.p.p.l
			if (lb && c == c.p.r) || (!lb && c == c.p.l) {
				// c is the middle of 4-node
				if c == c.p.r {
					c.p.leftRotate()
				} else {
					c.p.rightRotate()
				}
			} else {
				c = c.p
			}

			// rotate middle of 4-node to the root of subtree
			c.color = BLACK
			c.p.color = RED
			if lb {
				c.p.rightRotate()
			} else {
				c.p.leftRotate()
			}
		}
	}

	if c.p == Nil {
		if c.color != BLACK {
			c.color = BLACK
		}
		return c
	}

	return Nil
}

func (n *node) del() *node {
	// n is red, we can remove directly
	if n.color == RED {
		n.drop()
		return Nil
	} else if n.l.color == RED {
		n.l.color = BLACK
		n.drop()
		return Nil
	} else if n.r.color == RED {
		n.r.color = BLACK
		n.drop()
		return Nil
	}

	var x = n

	for x.p != Nil {
		// whether x is a left child
		lb := x.p.l == x
		s := x.sibling()

		if s.color == RED {
			// p is a 3-node
			x.p.color = RED
			s.color = BLACK
			if lb {
				x.p.leftRotate()
			} else {
				x.p.rightRotate()
			}
			s = x.sibling()
		}

		if s.l.color == BLACK && s.r.color == BLACK {
			// down overflow
			s.color = RED
			x = x.p
			if x.color == RED {
				x.color = BLACK
				break
			}
			continue
		}

		if lb {
			if s.r.color == BLACK {
				s.color = RED
				s.l.color = BLACK
				s.rightRotate()
				s = x.sibling()
			}
			s.color = x.p.color
			x.p.color = BLACK
			s.r.color = BLACK
			x.p.leftRotate()

		} else {
			if s.l.color == BLACK {
				s.color = RED
				s.r.color = BLACK
				s.leftRotate()
				s = x.sibling()
			}
			s.color = x.p.color
			x.p.color = BLACK
			s.l.color = BLACK
			x.p.rightRotate()
		}

		x = x.p
		break
	}

	n.drop()

	for x.p != Nil {
		x = x.p
	}

	return x
}

func (n *node) predecessor() *node {
	if n.l == Nil {
		return Nil
	}

	var c = Nil
	for c = n.l; c.r != Nil; c = c.r {
	}

	return c
}

func (n *node) successor() *node {
	if n.r == Nil {
		return Nil
	}

	var c = Nil
	for c = n.r; c.l != Nil; c = c.l {

	}
	return c
}

func (n *node) leftRotate() {
	/**
		          p                                    p
		          |                                    |
				  a                                    c
				/   \     a.leftRotate()             /   \
		       b     c   ================>		    a     g
	          / \   / \                            / \
		     d   e f   g                          b   f
		                                         / \
		                                        d   e
	*/

	a := n
	c := a.r

	a.r = c.l
	if c.l != Nil {
		c.l.p = a
	}
	c.l = a

	if a.p != Nil {
		if a == a.p.l {
			a.p.l = c
		} else {
			a.p.r = c
		}
	}

	c.p = a.p
	a.p = c

}

func (n *node) drop() {
	np := new(*node)
	if n.p.l == n {
		np = &n.p.l
	} else {
		np = &n.p.r
	}
	if n.l != Nil {
		*np = n.l
		n.l.p = n.p
	} else {
		*np = n.r
		n.r.p = n.p
	}
}

func (n *node) rightRotate() {
	/**
		          p                                    p
		          |                                    |
				  a                                    b
				/   \     a.leftRotate()             /   \
		       b     c   ================>		    d     a
	          / \   / \                                  / \
		     d   e f   g                                e   c
		                                                   / \
		                                                  f   g
	*/

	a := n
	b := a.l

	a.l = b.r
	if b.r != Nil {
		b.r.p = a
	}
	b.r = a

	if a.p != Nil {
		if a == a.p.l {
			a.p.l = b
		} else {
			a.p.r = b
		}
	}

	b.p = a.p
	a.p = b
}

func (n *node) grandparent() *node {
	if n.p != Nil {
		return n.p.p
	}
	return Nil
}

func (n *node) uncle() *node {
	gp := n.grandparent()
	if gp == Nil {
		return Nil
	}

	var u = Nil
	if gp.l == n.p {
		u = gp.r
	} else {
		u = gp.l
	}

	return u
}

func (n *node) sibling() *node {
	var s = Nil
	// n must has parent
	if n.p.l == n { // left is self
		s = n.p.r
	} else {
		s = n.p.l
	}

	return s
}

func (n *node) preOrder() {
	fmt.Println(n.string())
	if n.l != Nil {
		n.l.preOrder()
	}
	if n.r != Nil {
		n.r.preOrder()
	}
}

func (n *node) postOrder() {
	if n.l != Nil {
		n.l.postOrder()
	}
	if n.r != Nil {
		n.r.postOrder()
	}
	fmt.Println(n.string())
}

func (n *node) string() string {
	c := "r"
	if n.color == BLACK {
		c = "b"
	}

	return fmt.Sprintf("{color: %s, key: %v, val: %v}", c, n.k, n.v)
}

func New(cmp Compare) *Tree {
	return &Tree{
		cmp:  cmp,
		root: Nil,
	}
}

type Tree struct {
	root *node
	cmp  Compare
}

func (t *Tree) Ins(k interface{}, v interface{}) {
	n := &node{
		k:     k,
		v:     v,
		color: RED,
		p:     Nil,
		l:     Nil,
		r:     Nil,
	}

	if t.root == Nil {
		n.color = BLACK
		t.root = n
		return
	}

	np := new(*node)
	p := t.root
	// afterIns the new node
	for {
		if b := t.cmp(p.k, n.k); b == 0 {
			p.v = n.v
			return
		} else if b > 0 {
			np = &p.l
		} else {
			np = &p.r
		}
		if *np == Nil {
			*np = n
			n.p = p
			break
		} else {
			p = *np
		}
	}

	if r := n.afterIns(); r != Nil {
		t.root = r
	}
}

func (t *Tree) Del(key interface{}) interface{} {
	if t.root == Nil {
		return nil
	}

	var c = t.root
	for c != Nil {
		if b := t.cmp(c.k, key); b == 0 {
			break
		} else if b > 0 {
			c = c.l
		} else {
			c = c.r
		}
	}

	if c == Nil {
		return nil
	}

	v := c.v

	d := c.predecessor()
	if d == Nil {
		d = c.successor()
	}
	if d == Nil {
		d = c
	} else {
		c.k = d.k
		c.v = d.v
	}

	if r := d.del(); r != Nil {
		t.root = r
	}

	return v
}

func (t *Tree) Search(key interface{}) (interface{}, bool) {
	if t.root == Nil {
		return Nil, false
	}

	x := t.root
	for x != Nil {
		if b := t.cmp(x.k, key); b == 0 {
			return x.v, true
		} else if b > 0 {
			x = x.l
		} else {
			x = x.r
		}
	}

	return Nil, false

}

func (t *Tree) PreOder() {
	if t.root != Nil {
		t.root.preOrder()
	}
}

func (t *Tree) PostOrder() {
	if t.root != Nil {
		t.root.postOrder()
	}
}

func (t *Tree) IsRbTree() bool {
	if t.root == Nil {
		return true
	}

	if t.root.color != BLACK {
		return false
	}
	_, is := t.root.verify(0)
	return is
}

func (n *node) verify(h int) (int, bool) {
	if n == Nil {
		return h, true
	}

	if n.color == RED && n.p.color == RED {
		return 0, false
	}

	if n.color == BLACK {
		h++
	}

	lm, is := n.l.verify(h)
	if !is {
		return 0, false
	}
	rm, is := n.r.verify(h)
	if !is {
		return 0, false
	}
	return lm, lm == rm

}
