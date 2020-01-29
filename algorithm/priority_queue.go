package algorithm

import "leetcode/data_structure/heap"

type PriorityQueue struct {
	h *heap.KaryHeap
}

func NewPriorityQueue() *PriorityQueue {
	h := heap.NewKaryHeap(func(i, j interface{}) bool {
		return i.(Item).Val < j.(Item).Val
	}, heap.Kary(2))
	return &PriorityQueue{
		h: h,
	}
}

func (q *PriorityQueue) Push(key int, val float64) {
	idx := q.h.Walk(func(i interface{}) bool {
		if i.(Item).Key == key {
			return true
		}
		return false
	})
	if idx >= 0 {
		q.h.Remove(idx)
	}
	q.h.Push(Item{Key: key, Val: val})
}

func (q *PriorityQueue) Pop() (int, bool) {
	it := q.h.Pop()
	if it == nil {
		return 0, false
	}
	return it.(Item).Key, true
}

type Item struct {
	Key int
	Val float64
}
