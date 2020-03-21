package algorithm

import "unsafe"

/**
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/

// BST的中序遍历是升序序列
func isValidBSTByInOrder(root *TreeNode) bool {
	val := -1 << (8*unsafe.Sizeof(1) - 1)
	cur := root
	for cur != nil {
		if cur.Left == nil {
			if cur.Val > val {
				val = cur.Val
			} else {
				return false
			}
			cur = cur.Right
		} else {
			prev := cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}

			if prev.Right == nil {
				prev.Right = cur
				cur = cur.Left
			} else {
				if cur.Val > val {
					val = cur.Val
				} else {
					return false
				}
				prev.Right = nil
				cur = cur.Right
			}
		}
	}
	return true
}

// 通过BST的定义递归判断
func isValidBST(root *TreeNode) bool {
	min := -1 << (8*unsafe.Sizeof(1) - 1)
	max := 1<<(8*unsafe.Sizeof(1)-1) - 1
	return _isValidBST(root, min, max)
}

func _isValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	v := root.Val
	if root.Right != nil {
		if r := root.Right.Val; r <= min || r >= max || r <= v {
			return false
		}

	}
	if root.Left != nil {
		if l := root.Left.Val; l >= max || l <= min || l >= v {
			return false
		}
	}

	return _isValidBST(root.Right, v, max) && _isValidBST(root.Left, min, v)
}
