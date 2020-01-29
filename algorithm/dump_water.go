package algorithm

import (
	list2 "container/list"
	"fmt"
)

// 穷举法：首先要定义问题的解，并分析解空间的范围和拓扑结构，然后根据解空间的范围和拓扑结构设计遍历搜索算法
// 穷举法实际上就是搜索解空间树，搜索过程的推进实际为状态的转移，需要注意环的问题

const (
	BucketCount = 3
)

// capacity of bucket
var capacity = map[int]int{
	0: 8,
	1: 5,
	2: 3,
}

type Action struct {
	From  int
	To    int
	Water int
}

type State struct {
	bucket [BucketCount]int
	action Action
}

// 有三个没有刻度的桶，容量分别是8L，5L，3L，初始只有8L的水桶装满了水，现在需要借助其他两个水桶将其等分成两个4L
func DumpWater() {
	init := State{
		bucket: [BucketCount]int{8, 0, 0},
		action: Action{-1, 0, 8},
	}
	list := list2.New()
	list.PushBack(init)
	set := map[[BucketCount]int]bool{}
	set[init.bucket] = true
	searchSolution(list, set)
}

var actions = []Action{
	{
		From: 0,
		To:   1,
	},
	{
		From: 0,
		To:   2,
	},
	{
		From: 1,
		To:   2,
	},
	{
		From: 2,
		To:   1,
	},
	{
		From: 2,
		To:   0,
	},
	{
		From: 1,
		To:   0,
	},
}

func searchSolution(states *list2.List, filter map[[BucketCount]int]bool) {
	cur := states.Back()
	cs := cur.Value.(State)
	if finish(cs) {
		printSolution(states)
		return
	}

	for _, a := range actions {
		if canDump(cs, a.From, a.To) {
			// 计算可以从From倒多少水到To
			delta := min(cs.bucket[a.From], capacity[a.To]-cs.bucket[a.To])
			// 计算新的状态并判断是否该状态已经在路径中，避免出现环
			nb := cs.bucket
			nb[a.From] -= delta
			nb[a.To] += delta
			if filter[nb] {
				continue
			}

			filter[nb] = true
			states.PushBack(State{
				bucket: nb,
				action: Action{
					From:  a.From,
					To:    a.To,
					Water: delta,
				},
			})

			// dfs
			searchSolution(states, filter)
			delete(filter, nb)
			states.Remove(states.Back())
		}
	}

}

func canDump(state State, from, to int) bool {
	if state.bucket[from] == 0 || state.bucket[to] == capacity[to] {
		return false
	}
	return true
}

func finish(state State) bool {
	if state.bucket[0] == 4 && state.bucket[1] == 4 {
		return true
	}
	return false
}

func printSolution(states *list2.List) {
	s := states.Front()
	for s.Next() != nil {
		s = s.Next()
		v := s.Value.(State)
		fmt.Printf("from %d to %d with %dL; ", v.action.From, v.action.To, v.action.Water)
	}
	fmt.Println()
}
