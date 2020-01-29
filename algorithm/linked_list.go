package algorithm

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func arr2LinkedList(arr []int) *ListNode {
	head := new(ListNode)
	n := head
	for _, v := range arr {
		n.Next = &ListNode{
			Val: v,
		}
		n = n.Next
	}
	return head
}

func printList(tag string, head *ListNode) {
	fmt.Printf("%s[", tag)
	for n := head.Next; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
	fmt.Println("]")
}
