package algorithm

type listNode struct {
	Val  int
	Next *listNode
}

func detectCycle(head *listNode) *listNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == slow || fast == nil || slow == nil {
			break
		}
	}
	if slow == nil || fast == nil {
		return nil
	}
	p := slow
	q := head
	// 快指针是慢指针的2倍，假设前面慢指针走了N步，现在慢指针再走N步还会回到现在的位置（快指针已经验证过了），
	// 而从链头走N步也会到这个位置（慢指针已经验证过了），那么slow和head同步走N步，最终肯定会到达同一个位置
	// 那么他们肯定在链的入口就相遇了
	for p != q {
		p = p.Next
		q = q.Next
	}
	return p
}
