package algorithm

import (
	"fmt"
	"leetcode/data_structure/stack"
)

func TreeZigZag(t *Node) {
	if t == nil {
		return
	}
	s1 := stack.New()
	s2 := stack.New()
	rev := false
	s1.Push(t)
	for !s1.IsEmpty() {
		for !s1.IsEmpty() {
			n := s1.Pop().(*Node)
			fmt.Print(n.Value, " ")
			if rev {
				if n.right != nil {
					s2.Push(n.right)
				}
				if n.left != nil {
					s2.Push(n.left)
				}
			} else {
				if n.left != nil {
					s2.Push(n.left)
				}
				if n.right != nil {
					s2.Push(n.right)
				}
			}
		}
		s1, s2 = s2, s1
		rev = !rev
		fmt.Println()
	}

}
