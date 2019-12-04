package algorithm

import (
	"math"
	"testing"
)

func TestLongestSubstr(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcaadbc"))
}
func TestAddTwoNums(t *testing.T) {
	n1 := []int{3, 4, 2}
	n2 := []int{7}
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

func TestQuickSoft(t *testing.T) {
	arr := []int{2, 5, 6, 1, 7, 9, 0}
	QuickSoft(arr)
	t.Log(arr)
}

func TestReverseNodesInKGroup(t *testing.T) {
	n1 := []int{1, 2, 3}
	l1 := &ListNode{}
	cur := l1
	for _, v := range n1 {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	l1 = l1.Next
	l := reverseKGroup(l1, 2)

	for l != nil {
		t.Log(l.Val)
		l = l.Next
	}
}

func TestLongestValidParentheses(t *testing.T) {
	t.Log(longestValidParentheses("()(()())"))

	t.Log(longestValidParentheses("()())((((()(()))))()((()(()(())()))(()((()(())(((()((())())(((())(())())()()(()((((((()()(()())()))())()((()())((((((())()()()))(((()()((()()(()((((()))((()))(()(()())()(()((())())))(()()())()((((())"))
}

func TestLCS(t *testing.T) {
	t.Log(LCS("1253789", "1234789"))
}

func TestNumDistinct(t *testing.T) {
	t.Log(NumDistinct("rabbbit", "rabbit"))
	t.Log(NumDistinct("babgbag", "bag"))
}

func TestFindMedianFromDataStream(t *testing.T) {
	f := Constructor()
	f.AddNum(6)
	// 6
	t.Log(f.FindMedian()) // 6
	f.AddNum(10)
	// 6 10
	t.Log(f.FindMedian()) // 8
	f.AddNum(2)
	// 2 6 10
	t.Log(f.FindMedian()) //6
	f.AddNum(6)
	// 2 6 6 10

	t.Log(f.FindMedian()) // 6
	f.AddNum(5)
	// 2 5 6 6 10
	t.Log(f.FindMedian()) // 6
	// 2 5 6 6 6 10
	f.AddNum(6)

	f.AddNum(3)
	f.AddNum(1)
	t.Log(f.FindMedian())
}

func TestRob(t *testing.T) {
	ps, max := rob([]int{1, 7, 3, 1, 9})
	t.Log(ps, max)
}

func TestPrim(t *testing.T) {
	NO_EDGE := int64(math.MaxInt64)
	graph := [][]int64{
		{0, 1, NO_EDGE, 2, 3},
		{1, 0, 4, 6, NO_EDGE},
		{2, 4, 0, 3, 2},
		{NO_EDGE, 6, 3, 0, 7},
		{3, NO_EDGE, 2, 7, 0},
	}
	Prim(graph)
}

func TestCalc8Queue(t *testing.T) {
	queue_8()
}

func TestHanoi(t *testing.T) {
	Hanoi(4)
}

func TestEditDist(t *testing.T) {
	hello := []string{
		"hello", "helloo", "holle", "olle", "hello", "holl", "hell", "helo",
	}

	for i := range hello {
		t.Log(hello[i], editDist(hello[i], "hello"))
	}

	t.Log(editDist("xxxxyyyyxx", "yyyyy"))

}

func TestCalculateMinimumHP(t *testing.T) {
	t.Log(calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}))
	t.Log(calculateMinimumHP([][]int{
		{1, -2, 3},
		{2, -2, -2},
	}))
	t.Log(calculateMinimumHP([][]int{{-3, 5}}))
	t.Log(calculateMinimumHP([][]int{{0}}))
	t.Log(calculateMinimumHP([][]int{{200}}))
	t.Log(calculateMinimumHP([][]int{{-200}}))
}

func TestThrowEgg(t *testing.T) {
	t.Log(throwEgg(100, 2)) // 14
	t.Log(throwEgg(200, 2)) // 20
	t.Log(throwEgg(10, 3))
}
