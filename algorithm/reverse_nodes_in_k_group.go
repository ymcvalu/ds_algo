package algorithm

/**
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	var slow, fast = head, head
	var _head, tail *ListNode

	for slow != nil {
		for i := 0; i < k && fast != nil; i++ {
			fast = fast.Next
		}

		cur := slow
		next := cur.Next
		for next != fast {
			next.Next, next, cur = cur, next.Next, next
		}

		if _head == nil {
			_head = cur
		} else {
			tail.Next = cur
		}

		tail = slow
		slow = fast
	}
	tail.Next=nil

	return _head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	var h, t *ListNode

	for head != nil {
		pre := head

		for i := 1; i < k; i++ {
			pre = pre.Next
			if pre == nil {
				break
			}
		}

		if pre == nil {
			if h == nil {
				h = head
			} else {
				t.Next = head
			}
			break
		}

		pre = head
		post := head.Next
		for i := 1; post != nil && i < k; i++ {
			cur := post
			post = post.Next
			cur.Next = pre
			pre = cur
		}

		if h == nil {
			h = pre
		} else {
			t.Next = pre
		}
		t = head
		t.Next = nil
		head = post
	}

	return h
}
