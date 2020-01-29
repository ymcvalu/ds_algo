package algorithm

import (
	"container/list"
	"leetcode/data_structure/graph"
)

/**
Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:

Only one letter can be changed at a time
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.


Note:

Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.

Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]

Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: []

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

*/

// 思路：本质上是图的最短路径搜索，可用假设所有路径权重都是1，使用bfs就可以了

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	lw := len(wordList)
	if lw == 0 {
		return nil
	}

	// 构造图
	g := graph.New(lw+1, false)

	endV := -1
	for i := 0; i < lw; i++ {
		if wordList[i] == endWord {
			endV = i + 1
			break
		}
	}

	if endV == -1 {
		return nil
	}

	for i := 0; i < lw; i++ {
		if beginWord != wordList[i] && connected(beginWord, wordList[i]) {
			g.AddEdge(0, i+1)
		}

		for j := i + 1; j < lw; j++ {
			if connected(wordList[i], wordList[j]) {
				g.AddEdge(i+1, j+1)
			}
		}
	}

	// 不可达
	if len(g.Adj(endV)) == 0 {
		return nil
	}

	dist := make([]int, lw+1)
	marked := make(map[int]bool, lw+1)
	marked[0] = true
	l := list.New()
	l.PushBack(0)
	// BFS
	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		v := e.Value.(int)
		for _, n := range g.Adj(v) {
			if marked[n] {
				continue
			}
			dist[n] = dist[v] + 1
			marked[n] = true

			if n == endV {
				continue
			}

			l.PushBack(n)
		}
	}

	// 不可达
	if dist[endV] == 0 {
		return nil
	}

	// 回溯所有最短路径
	var results [][]string
	var trace func(int, int, []string)
	trace = func(v, n int, result []string) {
		if n == 0 {
			// 需要深拷贝
			r := make([]string, len(result))
			copy(r, result)
			results = append(results, r)
			return
		}

		for _, adj := range g.Adj(v) {
			if dist[adj] == n {
				result[n] = wordList[adj-1]
				trace(adj, n-1, result)
			}
		}
	}

	result := make([]string, dist[endV]+1)
	result[dist[endV]] = endWord
	result[0] = beginWord
	trace(endV, dist[endV]-1, result)

	return results
}

func connected(from, to string) bool {
	n := 0
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			n++
			if n > 1 {
				return false
			}
		}
	}
	return true
}
