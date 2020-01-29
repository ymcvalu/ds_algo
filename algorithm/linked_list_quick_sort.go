package algorithm

// head为哨兵节点
func LinkedListQuickSort(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	lh, mh, rh := linkedListPartition(head)

	LinkedListQuickSort(lh)
	LinkedListQuickSort(rh)

	m := mh
	for m.Next != nil {
		m = m.Next
	}
	m.Next = rh.Next

	l := lh
	for l.Next != nil {
		l = l.Next
	}
	l.Next = mh.Next

	head.Next = lh.Next

	return
}

// head为哨兵节点
func linkedListPartition(head *ListNode) (lh, mh, rh *ListNode) {
	lh = new(ListNode)
	mh = new(ListNode)
	rh = new(ListNode)

	v := head.Next.Val

	n := head.Next.Next

	mh.Next = head.Next
	mh.Next.Next = nil

	for n != nil {
		_n := n
		n = _n.Next
		if _n.Val == v {
			_n.Next = mh.Next
			mh.Next = _n
		} else if _n.Val > v {
			_n.Next = rh.Next
			rh.Next = _n
		} else {
			_n.Next = lh.Next
			lh.Next = _n
		}
	}
	return
}

