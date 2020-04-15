package algorithm

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return _isSymmetric(root, root)
}

func _isSymmetric(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil || n1.Val != n2.Val {
		return false
	}
	return _isSymmetric(n1.Left, n2.Right) && _isSymmetric(n1.Right, n2.Left)
}
