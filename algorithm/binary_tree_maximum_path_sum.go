package algorithm

import (
	"unsafe"
)

/**
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections.
The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6

Example 2:

Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
*/

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max = -1 << (8*unsafe.Sizeof(int(1)) - 1)
	_maxPathSum(root, &max)
	return max
}

func _maxPathSum(n *TreeNode, max *int) int {
	if n == nil {
		return 0
	}
	lm := _maxPathSum(n.Left, max)
	rm := _maxPathSum(n.Right, max)

	if lm < 0 {
		lm = 0
	}

	if rm < 0 {
		rm = 0
	}

	if max == nil || lm+rm+n.Val > *max {
		*max = lm + rm + n.Val
	}

	if lm > rm {
		return lm + n.Val
	}
	return rm + n.Val
}
