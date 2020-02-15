package algorithm

import "fmt"

// morris遍历，实际上是构造线索树，利用线索返回父节点，当第二次返回父节点时再移除线索
// 是一种时间换空间的遍历方法

// 时间复杂度O(n) 空间复杂度O(1)
func morrisTraversePreOrder(t *Node) {
	cur := t
	for cur != nil {
		if cur.left == nil {
			fmt.Print(cur.Value, " ")
			cur = cur.right
		} else {
			prev := cur.left
			// 查找prev的平均时机复杂度为O(log n)，但是所有节点总的时机复杂度为O(n)
			for prev.right != nil && prev.right != cur {
				prev = prev.right
			}

			if prev.right == nil {
				fmt.Print(cur.Value, " ")
				prev.right = cur
				cur = cur.left
			} else {
				prev.right = nil
				cur = cur.right
			}
		}
	}
}

// 时间复杂度O(n) 空间复杂度O(1)
func morrisTraverseInOrder(t *Node) {
	cur := t
	for cur != nil {
		if cur.left == nil {
			fmt.Print(cur.Value, " ")
			cur = cur.right
		} else {
			prev := cur.left
			for prev.right != nil && prev.right != cur {
				prev = prev.right
			}

			if prev.right == nil {
				prev.right = cur
				cur = cur.left
			} else {
				fmt.Print(cur.Value, " ")
				prev.right = nil
				cur = cur.right
			}
		}
	}
}

// 时间复杂度O(n) 空间复杂度O(1)
func morrisTraversePostOrder(t *Node) {
	if t == nil {
		return
	}

	cur := &Node{
		left: t,
	}

	for cur != nil {
		if cur.left == nil {
			cur = cur.right
		} else {
			prev := cur.left
			for prev.right != nil && prev.right != cur {
				prev = prev.right
			}
			if prev.right == nil {
				prev.right = cur
				cur = cur.left
			} else {
				prev.right = nil
				printRightEdgeReverse(cur.left)
				cur = cur.right
			}

		}

	}
}

func printRightEdgeReverse(n *Node) {
	var (
		cur        = n
		next, prev *Node
	)
	for cur != nil {
		next = cur.right
		cur.right = prev
		prev, cur = cur, next
	}

	cur = prev
	for c := cur; c != nil; c = c.right {
		fmt.Print(c.Value, " ")
	}

	prev = nil
	for cur != nil {
		next = cur.right
		cur.right = prev
		prev, cur = cur, next
	}

}
