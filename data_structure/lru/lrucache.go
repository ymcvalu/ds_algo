package lru

import (
	"time"
	"sync"
	"container/list"
)

type Elem struct {
	Key          string
	Value        interface{}
	timeAccessed time.Time
}

type LRUCache struct {
	sync.Mutex
	list        *list.List
	elems       map[string]*list.Element
	maxlifetime int64
}

func New(maxlifetime int64) *LRUCache {
	m := &LRUCache{
		maxlifetime: maxlifetime,
		list:        list.New(),
		elems:       make(map[string]*list.Element),
	}
	go func() {

		c := time.Tick(time.Second * 1)
		select {
		case <-c:
			m.gc()
		}
	}()
	return m
}

func (l *LRUCache) Set(key string, value interface{}) {
	l.Lock()
	defer l.Unlock()
	if elem, ok := l.elems[key]; ok {
		l.list.Remove(elem)
	}
	e := &Elem{
		Key:          key,
		Value:        value,
		timeAccessed: time.Now(),
	}
	elem := l.list.PushFront(e)
	l.elems[key] = elem
}

func (l *LRUCache) Get(key string) (interface{}, bool) {
	l.Lock()
	defer l.Unlock()
	if elem, ok := l.elems[key]; ok {
		e := elem.Value.(*Elem)
		e.timeAccessed = time.Now()
		l.list.MoveToFront(elem)
		return e.Value, true
	} else {
		return nil, false
	}
}

func (l *LRUCache) Del(key string) {
	l.Lock()
	defer l.Unlock()
	elem := l.elems[key]
	l.list.Remove(elem)
	delete(l.elems, key)
}

func (l *LRUCache) gc() {
	l.Lock()
	defer l.Unlock()
	for {
		e := l.list.Back()

		if e == nil {
			break
		}

		elem := e.Value.(*Elem)
		if elem.timeAccessed.Unix()+l.maxlifetime > time.Now().Unix() {
			break
		}
		l.list.Remove(e)
		delete(l.elems, elem.Key)
	}

}
