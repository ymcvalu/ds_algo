package algorithm

// 二叉树的任意两节点的最近公共祖先，假设节点值唯一

func LCA(root *Node, v1, v2 int) *Node {
	set := make(map[*Node]uint8)
	_, p := lca(root, v1, v2, set)
	return p
}

func lca(n *Node, v1, v2 int, set map[*Node]uint8) (bool, *Node) {
	if n == nil {
		return false, nil
	}

	has := false
	if n.Value == v1 || n.Value == v2 {
		has = true
		set[n] = set[n] + 1
	}

	chs := make([]*Node, 2)
	chs[0] = n.left
	chs[1] = n.right

	for _, c := range chs {
		is, p := lca(c, v1, v2, set)
		if p != nil {
			return true, p
		} else if is {
			has = true
			set[n] = set[n] + 1
			if set[n] == 2 {
				return true, n
			}
		}
	}

	return has, nil
}
