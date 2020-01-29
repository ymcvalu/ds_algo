package algorithm

import "container/list"

type Stack struct {
	q1 *list.List
	q2 *list.List
}

func NewStack() *Stack {
	return &Stack{
		q1: list.New(),
		q2: list.New(),
	}
}

func (s *Stack) Push(i int) {
	s.q1.PushBack(i)
}

func (s *Stack) Pop() (int, bool) {
	for s.q1.Len() > 1 {
		e := s.q1.Front()
		s.q2.PushBack(e.Value)
		s.q1.Remove(e)
	}

	if s.q1.Len() == 0 {
		return 0, false
	}

	e := s.q1.Front()
	s.q1.Remove(e)
	s.q1, s.q2 = s.q2, s.q1
	return e.Value.(int), true
}
