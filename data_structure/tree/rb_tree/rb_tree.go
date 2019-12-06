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

func (n *node) afterIns() *node {
	c := n
	for c.p != nil && c.isRed() && c.p.isRed() {
		if u := c.uncle(); u.isRed() {
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

	if c.p == nil {
		if c.isRed() {
			c.color = BLACK
		}
		return c
	}

	return nil
}

func (n *node) del() *node {
	// n is red, we can remove directly
	if n.isRed() {
		n.drop()
		return nil
	} else if n.l.isRed() {
		n.l.color = BLACK
		n.drop()
		return nil
	} else if n.r.isRed() {
		n.r.color = BLACK
		n.drop()
		return nil
	}

	var x = n

	for x.p != nil {
		// whether x is a left child
		lb := x.p.l == x
		s := x.sibling()

		if s.isRed() {
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

		if s.l.isBlack() && s.r.isBlack() {
			// down overflow
			s.color = RED
			x = x.p
			if x.isRed() {
				x.color = BLACK
				break
			}
			continue
		}

		if lb {
			if s.r.isBlack() {
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
			if s.l.isBlack() {
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

	for x.p != nil {
		x = x.p
	}

	return x
}

func (n *node) predecessor() *node {
	if n.l == nil {
		return nil
	}

	var c *node
	for c = n.l; c.r != nil; c = c.r {
	}

	return c
}

func (n *node) successor() *node {
	if n.r == nil {
		return nil
	}

	var c *node
	for c = n.r; c.l != nil; c = c.l {

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
	if c.l != nil {
		c.l.p = a
	}
	c.l = a

	if a.p != nil {
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
	if n.l != nil {
		*np = n.l
		if n.l != nil {
			n.l.p = n.p
		}
	} else {
		*np = n.r
		if n.r != nil {
			n.r.p = n.p
		}
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
	if b.r != nil {
		b.r.p = a
	}
	b.r = a

	if a.p != nil {
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
	if n.p != nil {
		return n.p.p
	}
	return nil
}

func (n *node) uncle() *node {
	gp := n.grandparent()
	if gp == nil {
		return nil
	}

	var u *node
	if gp.l == n.p {
		u = gp.r
	} else {
		u = gp.l
	}

	return u
}

func (n *node) sibling() *node {
	var s *node
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
	if n.l != nil {
		n.l.preOrder()
	}
	if n.r != nil {
		n.r.preOrder()
	}
}

func (n *node) postOrder() {
	if n.l != nil {
		n.l.postOrder()
	}
	if n.r != nil {
		n.r.postOrder()
	}
	fmt.Println(n.string())
}

func (n *node) isRed() bool {
	return n != nil && n.color == RED
}

func (n *node) isBlack() bool {
	return n == nil || n.color == BLACK
}

func (n *node) string() string {
	c := "r"
	if n.isBlack() {
		c = "b"
	}

	return fmt.Sprintf("{color: %s, key: %v, val: %v}", c, n.k, n.v)
}

func New(cmp Compare) *Tree {
	return &Tree{
		cmp:  cmp,
		root: nil,
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
		p:     nil,
		l:     nil,
		r:     nil,
	}

	if t.root == nil {
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
		if *np == nil {
			*np = n
			n.p = p
			break
		} else {
			p = *np
		}
	}

	if r := n.afterIns(); r != nil {
		t.root = r
	}
}

func (t *Tree) Del(key interface{}) interface{} {
	if t.root == nil {
		return nil
	}

	var c = t.root
	for c != nil {
		if b := t.cmp(c.k, key); b == 0 {
			break
		} else if b > 0 {
			c = c.l
		} else {
			c = c.r
		}
	}

	if c == nil {
		return nil
	}

	v := c.v

	d := c.predecessor()
	if d == nil {
		d = c.successor()
	}
	if d == nil {
		d = c
	} else {
		c.k = d.k
		c.v = d.v
	}

	if r := d.del(); r != nil {
		t.root = r
	}

	return v
}

func (t *Tree) Search(key interface{}) (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}

	x := t.root
	for x != nil {
		if b := t.cmp(x.k, key); b == 0 {
			return x.v, true
		} else if b > 0 {
			x = x.l
		} else {
			x = x.r
		}
	}

	return nil, false

}

func (t *Tree) PreOder() {
	if t.root != nil {
		t.root.preOrder()
	}
}

func (t *Tree) PostOrder() {
	if t.root != nil {
		t.root.postOrder()
	}
}

func (t *Tree) IsRbTree() bool {
	if t.root == nil {
		return true
	}

	if t.root.isRed() {
		return false
	}
	_, is := t.root.verify(0)
	return is
}

func (n *node) verify(h int) (int, bool) {
	if n == nil {
		return h, true
	}

	if n.isRed() && n.p.isRed() {
		return 0, false
	}

	if n.isBlack() {
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
