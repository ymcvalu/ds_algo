package algorithm

/**
给定一个带有头结点 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。
*/
func middleNode(head *ListNode) *ListNode {
	// 使用快慢指针
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}

		fast = fast.Next
		if fast == nil {
			break
		}
	}
	return slow
}
