package algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbersReverse(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseLinkedList(l1)
	l2 = reverseLinkedList(l2)

	head := &ListNode{}
	cur := head
	overflow := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + overflow
		overflow = sum / 10
		sum %= 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 == nil {
		l1 = l2
	}

	if overflow != 0 {
		for overflow != 0 && l1 != nil {
			sum := overflow + l1.Val
			overflow = sum / 10
			sum %= 10
			cur.Next = &ListNode{Val: sum}
			cur = cur.Next
			l1 = l1.Next
		}
		if overflow != 0 {
			cur.Next = &ListNode{Val: overflow}
			overflow = 0
		}
	}
	if l1 != nil {
		cur.Next = l1
	}

	return reverseLinkedList(head.Next)
}

func reverseLinkedList(l *ListNode) *ListNode {
	if l == nil {
		return l
	}
	pre, cur := l, l.Next
	pre.Next = nil
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
