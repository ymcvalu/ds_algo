package algorithm

import "github.com/ymcvalu/ds_algo/data_structure/stack"

type Queue struct {
	stack1 *stack.Stack
	stack2 *stack.Stack
}

func (q *Queue) Push(v int) {
	q.stack1.Push(v)
}

func (q *Queue) Pop() (int, bool) {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			q.stack2.Push(q.stack1.Pop())
		}
	}
	if q.stack2.IsEmpty() {
		return 0, false
	}
	return q.stack2.Pop().(int), true
}

func NewQueue() *Queue {
	return &Queue{
		stack1: stack.New(),
		stack2: stack.New(),
	}
}
