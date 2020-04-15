package algorithm

func TreeToList(root *Node) *ListNode {
	if root == nil {
		return nil
	}
	h := &ListNode{}
	treeToList(root, h, nil)
	return h
}

func treeToList(tn *Node, head *ListNode, pre *ListNode) *ListNode {
	if tn.left != nil {
		pre = treeToList(tn.left, head, pre)
	}
	if pre != nil {
		pre.Next = &ListNode{
			Val: tn.Value,
		}
		pre = pre.Next
	} else {
		head.Next = &ListNode{Val: tn.Value}
		pre = head.Next
	}
	if tn.right!=nil{
		pre=treeToList(tn.right,head,pre )
	}
	return pre
}
