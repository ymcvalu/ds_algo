package algorithm

import (
	"testing"
)

func TestLongestSubstr(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcaadbc"))
}
func TestAddTwoNums(t *testing.T) {
	n1 := []int{3, 4, 2}
	n2 := []int{7,}
	l1 := &ListNode{}
	cur := l1
	for _, v := range n1 {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	l2 := &ListNode{}
	cur = l2
	for _, v := range n2 {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	l3 := addTwoNumbers(l1.Next, l2.Next)
	for cur := l3; cur != nil; cur = cur.Next {
		t.Log(cur.Val)
	}
}

func TestMedianNum(t *testing.T) {
	num1 := []int{1}
	num2 := []int{2, 3}
	t.Log(findMedianSortedArrays(num1, num2))
}

func TestLongestPalindromicSubstring(t *testing.T) {
	t.Log(longestPalindrome("abccb"))
}

func TestRegMatch(t *testing.T) {
	t.Log(isMatch("bbab", "b*a*b"))
	t.Log(isMatch("bbbba", ".*a*a"))
	t.Log(isMatch("mississippi", "mis*is*p*.*"))
}

func TestZigZagConvert(t *testing.T) {
	t.Log(convert("AB", 1))
}
