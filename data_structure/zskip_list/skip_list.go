package zskip_list

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// redis版本的skipList

const (
	ZSKIP_LIST_MAXLEVEL = 64
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type znode struct {
	key      []byte
	score    float64
	backward *znode
	span     []int
	level    []*znode
}

func createNode(level int, score float64, key []byte) *znode {
	return &znode{
		key:   key,
		score: score,
		level: make([]*znode, level),
		span:  make([]int, level),
	}
}

type SkipList struct {
	header *znode
	tail   *znode
	length int
	level  int
}

func New() *SkipList {
	return &SkipList{
		header: createNode(ZSKIP_LIST_MAXLEVEL, 0, nil),
		level:  1,
	}
}

// caller需要保证key唯一
func (sl *SkipList) Insert(score float64, key []byte) {
	if math.IsNaN(score) {
		panic("unsupported: the score is NaN")
	}

	update := make([]*znode, ZSKIP_LIST_MAXLEVEL)
	rank := make([]int, ZSKIP_LIST_MAXLEVEL)
	x := sl.header

	// TODO: 这块代码提取出来做公共逻辑
	for i := sl.level - 1; i >= 0; i-- {
		if i < sl.level-1 {
			rank[i] = rank[i+1]
		}

		for x.level[i] != nil &&
			// 首先比较分数
			(x.level[i].score < score ||
				// 分数相同则按照key的字节序排序
				(x.level[i].score == score && bytes.Compare(x.level[i].key, key) < 0)) {
			rank[i] += x.span[i]
			x = x.level[i]
		}
		update[i] = x
	}

	level := randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			rank[i] = 0
			update[i] = sl.header
			update[i].span[i] = sl.length
		}
		sl.level = level
	}

	x = createNode(level, score, key)
	for i := 0; i < level; i++ {
		x.level[i] = update[i].level[i]
		update[i].level[i] = x
		// rank[0]-rank[i]表示从update[i]到当前节点的span
		x.span[i] = update[i].span[i] - (rank[0] - rank[i])
		update[i].span[i] = (rank[0] - rank[i]) + 1
	}

	for i := level; i < sl.level; i++ {
		update[i].span[i]++
	}

	if update[0] != sl.header {
		x.backward = update[0]
	}

	if x.level[0] != nil {
		x.level[0].backward = x
	} else {
		sl.tail = x
	}
	sl.length++
}

func (sl *SkipList) Delete(score float64, key []byte) bool {
	updates := make([]*znode, ZSKIP_LIST_MAXLEVEL)

	x := sl.header

	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i] != nil &&
			(x.level[i].score < score ||
				(x.level[i].score == score && bytes.Compare(x.level[i].key, key) < 0)) {
			x = x.level[i]
		}
		updates[i] = x
	}

	x = x.level[0]
	if x.score != score || bytes.Compare(x.key, key) != 0 {
		return false
	}

	sl.delNode(x, updates)

	return true
}

//  0 获取第一个元素
// -1 获取最后一个元素
func (sl *SkipList) GetByRank(r int) (float64, []byte, bool) {
	if r < 0 {
		r = int(sl.length) + r
	}

	if r < 0 || r >= sl.length {
		return 0, nil, false
	}

	r += 1 // skip header

	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i] != nil && r-x.span[i] >= 0 {
			r -= x.span[i]
			x = x.level[i]
		}
	}

	if r == 0 {
		return x.score, x.key, true
	}

	return 0.0, nil, false
}

func (sl *SkipList) delNode(x *znode, updates []*znode) {
	for i := 0; i < sl.level; i++ {
		if updates[i].level[i] == x {
			updates[i].span[i] += x.span[i] - 1
			updates[i].level[i] = x.level[i]
		} else {
			updates[i].span[i]--
		}
	}

	if x.level[0] != nil {
		x.level[0].backward = x.backward
	} else {
		sl.tail = x.backward
	}

	for sl.level > 1 && sl.header.level[sl.level-1] == nil {
		sl.level--
	}

	sl.length--
}

func (sl *SkipList) Length() int {
	return sl.length
}

func randomLevel() int {
	h := 1

	for h < ZSKIP_LIST_MAXLEVEL && rand.Int()&0xFFFF < 0xFFFF>>2 {
		h++
	}

	return h
}

// help debug
func (sl *SkipList) print() {
	c := sl.header
	for c != nil {
		span := make([]int, 0, sl.level)
		for i := 0; i < len(c.level) && i < sl.level; i++ {
			span = append(span, c.span[i])
		}

		fmt.Printf("score => %f ][ key => %s ][ span => %v\n", c.score, string(c.key), span)
		c = c.level[0]
	}
}

func (sl *SkipList) print_rev() {
	c := sl.tail
	for c != nil {
		fmt.Printf("score = > %f ][ key => %s\n", c.score, string(c.key))
		c = c.backward
	}
}
