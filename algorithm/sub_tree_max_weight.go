package algorithm

func GetMaxWeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0

	_getMaxWeight(root, &max)
	return max
}

func _getMaxWeight(n *TreeNode, max *int) int {
	var (
		lm int
		rm int
	)

	// 不能包含3度节点
	// TODO: root没有父节点，最多是2度节点
	if n.Left != nil && n.Right != nil {
		_lm := GetMaxWeight(n.Left)
		_rm := GetMaxWeight(n.Right)
		if _lm > _rm && _lm > *max {
			*max = _lm
		} else if _lm < _rm && _rm > *max {
			*max = _rm
		}
		return 0
	}

	if n.Left != nil {
		lm = _getMaxWeight(n.Left, max)

	}

	if n.Right != nil {
		rm = _getMaxWeight(n.Right, max)
	}

	if lm+rm+n.Val > *max {
		*max = lm + rm
	}

	return lm + rm + n.Val

}
