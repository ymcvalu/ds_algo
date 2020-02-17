package skip_list

import (
	"errors"
	"math/rand"
	"time"
)

var (
	ErrNotFound   = errors.New("key not found")
	ErrHasExisted = errors.New("key has existed")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Compare func(i, j interface{}) int

type Node struct {
	key  interface{}
	val  interface{}
	next []*Node
}

type SkipList struct {
	head     Node
	maxLevel int
	cmp      Compare
}

func New(maxLevel int, cmp Compare) *SkipList {
	sl := &SkipList{
		maxLevel: maxLevel,
		cmp:      cmp,
		head: Node{
			next: make([]*Node, maxLevel),
		},
	}

	return sl
}

func (s *SkipList) findGreaterOrEqual(key interface{}, path []*Node) *Node {
	x := &s.head
	level := len(x.next) - 1
	for {
		if next := x.next[level]; next != nil && s.cmp(key, next.key) > 0 {
			x = next
		} else {
			if path != nil {
				path[level] = x
			}

			if level == 0 {
				return next
			}
			level--
		}
	}

	return x
}

func (s *SkipList) Insert(key interface{}, val interface{}) error {
	path := make([]*Node, s.maxLevel)
	x := s.findGreaterOrEqual(key, path)
	if x != nil && s.cmp(x.key, key) == 0 {
		return ErrHasExisted
	}

	h := s.randomHeight()
	n := &Node{
		key:  key,
		val:  val,
		next: make([]*Node, h),
	}

	for i := 0; i < h; i++ {
		n.next[i] = path[i].next[i]
		path[i].next[i] = n
	}

	return nil
}

func (s *SkipList) Get(key interface{}) (interface{}, bool) {
	x := s.findGreaterOrEqual(key, nil)

	if x != nil && s.cmp(x.key, key) == 0 {
		return x.val, true
	}
	return nil, false
}

func (s *SkipList) Delete(key interface{}) error {
	path := make([]*Node, s.maxLevel)
	x := s.findGreaterOrEqual(key, path)
	if x == nil || s.cmp(x.key, key) != 0 {
		return ErrNotFound
	}

	h := len(x.next)
	for i := 0; i < h; i++ {
		path[i].next[i] = x.next[i]
	}

	return nil
}

func (s *SkipList) randomHeight() int {
	h := 1

	for h < s.maxLevel && rand.Int()&0xFFFF < 0xFFFF>>2 {
		h++
	}

	return h
}
