package stack

import "container/list"

type Stack struct {
	list *list.List
}

func New() *Stack {
	return &Stack{
		list: list.New(),
	}
}


func (s *Stack) Push(v interface{}) {
	s.list.PushFront(v)
}

func (s *Stack) Pop() interface{} {
	v := s.list.Front()
	if v != nil {
		s.list.Remove(v)
		return v.Value
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	v := s.list.Front()
	if v != nil {
		return v.Value
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
