package algorithm

import (
	list2 "container/list"
	"fmt"
)

func TraverseBinaryTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Println("preOrder")
	PreOrder(root)
	fmt.Println("\ninOrder")
	InOrder(root)
	fmt.Println("\npostOrder")
	PostOrder(root)
	fmt.Println("\nbfs")
	Bfs(root)
	fmt.Println()
}

func PreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.Value, " ")
	PreOrder(n.left)
	PreOrder(n.right)
}

func InOrder(n *Node) {
	if n == nil {
		return
	}

	InOrder(n.left)
	fmt.Print(n.Value, " ")
	InOrder(n.right)
}

func PostOrder(n *Node) {
	if n == nil {
		return
	}
	PostOrder(n.left)
	PostOrder(n.right)
	fmt.Print(n.Value, " ")
}

func Bfs(root *Node) {
	if root == nil {
		return
	}
	l := list2.New()
	l.PushBack(root)

	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		n := e.Value.(*Node)
		fmt.Print(n.Value, " ")
		if n.left != nil {
			l.PushBack(n.left)
		}
		if n.right != nil {
			l.PushBack(n.right)
		}
	}
}
