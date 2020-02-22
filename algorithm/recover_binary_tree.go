package algorithm

// 根据前序遍历和中序遍历结果恢复BT（保证不包含重复数字）
func RecoverBinaryTree(pre, in []int) *Node {
	if len(pre) != len(in) {
		panic("invalid parameters")
	}
	if len(pre) == 0 {
		return nil
	}

	n := &Node{
		Value: pre[0],
	}

	i := 0
	for {
		if in[i] == n.Value {
			break
		}
		i++
	}

	n.left = RecoverBinaryTree(pre[1:i+1], in[:i])
	n.right = RecoverBinaryTree(pre[i+1:], in[i+1:])
	return n
}