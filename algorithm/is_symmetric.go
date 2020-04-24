package algorithm

func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}

	return _isSymmetric(root, root)
}

func _isSymmetric(n1, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 == nil || n2 == nil || n1.Value != n2.Value {
		return false
	}
	return _isSymmetric(n1.left, n2.right) && _isSymmetric(n1.right, n2.left)
}
