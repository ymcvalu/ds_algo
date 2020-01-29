package algorithm

import "container/list"

type LinkedList struct {
	l *list.List
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		l: list.New(),
	}
}

func (l *LinkedList) Push(v interface{}) {
	l.l.PushBack(v)
}

func (l *LinkedList) Pop() interface{} {
	e := l.l.Front()
	if e != nil {
		l.l.Remove(e)
	}
	return e.Value
}

func (l *LinkedList) IsEmpty() bool {
	return l.l.Len() == 0
}
