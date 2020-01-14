package algorithm

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestLongestSubstr(t *testing.T) {
	t.Log(lengthOfLongestSubstring1("abcaadbcefsad"))
	t.Log(lengthOfLongestSubstring2("abcaadbcefsad"))
	t.Log(lengthOfLongestSubstring2("abcdefg"))
	t.Log(lengthOfLongestSubstring2("aaabedks"))
	t.Log(lengthOfLongestSubstring2("adbcdd"))
	t.Log(lengthOfLongestSubstring2("aadfghjkllkjqwertyuiioppoi"))
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
	t.Log(longestPalindrome1("abccb"), longestPalindrome2("abccb"))
	t.Log(longestPalindrome1("abccbajklsdfjkabbbbbajskldfjlk"), longestPalindrome2("abccbajklsdfjkabbbbbajskldfjlk"))
	t.Log(longestPalindrome1("sdfdsfdsfjjddkkddjjksl"), longestPalindrome2("sdfdsfdsfjjddkkddjjksl"))
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
	t.Log(calculateMinimumHP([][]int{{0, 0, 0}, {1, 1, -1}}))
	t.Log(calculateMinimumHP([][]int{{1, 2, 1}, {-2, -3, -3}, {3, 2, -2}}))
	t.Log(calculateMinimumHP([][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}))
}

func TestThrowEgg(t *testing.T) {
	t.Log(throwEgg(100, 2)) // 14
	t.Log(throwEgg(200, 2)) // 20
	t.Log(throwEgg(10, 3))
}

func TestFindOneElem3(t *testing.T) {
	t.Log(findOneElem3([]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 4}))
}

func TestFindOneElem2(t *testing.T) {
	t.Log(findOneElem2([]int{1, 2, 1, 2, 3}))
}

func TestSwapLetter(t *testing.T) {
	t.Log(swapLetter("abcbaa", 2))
	t.Log(swapLetter("abcbaa", 3))
	t.Log(swapLetter("dabdcddsd", 4))
	t.Log(swapLetter("dabdcddsd", 5))

}

func TestSkipStage(t *testing.T) {
	t.Log(skipStage(1))
	t.Log(skipStage(2))
	t.Log(skipStage(3))
	t.Log(skipStage(4))
	t.Log(skipStage(5))
	t.Log(skipStage(6))
	t.Log(skipStage(7))
}

func TestMaxArea(t *testing.T) {
	t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func TestSidlingWindowMaximum(t *testing.T) {
	t.Log(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7, 5, 4, 6, 9, 7, 2}, 3))
}

func TestMaxProfit(t *testing.T) {
	t.Log(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	t.Log(maxProfit([]int{7, 6, 4, 3, 1}))

	t.Log(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	t.Log(maxProfit2([]int{1, 2, 3, 4, 5}))

	t.Log(maxProfit3([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	t.Log(maxProfit3([]int{1, 3, 0, 5, 1, 2, 1, 7}))
	t.Log(maxProfit3([]int{1, 2, 3, 4, 5}))
	t.Log(maxProfit3([]int{7, 6, 4, 3, 1}))
	t.Log(maxProfit3([]int{0, 3, 0, 5, 0, 2}))

}

func TestIsPalindrome(t *testing.T) {
	t.Log(isPalindrome("A man, a plan, a canal: Panama"))
	t.Log(isPalindrome("race a car`"))
	t.Log(isPalindrome(".,"))
}

func TestReverseInteger(t *testing.T) {
	t.Log(reverse(123), reverse(-123), reverse(120), reverse(102), reverse(2147483647), reverse(-2147483412))
}

func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("32"), myAtoi("-32"), myAtoi("+32"), myAtoi("+0032"), myAtoi("-0032"))
	t.Log(myAtoi("  +32"), myAtoi("   -32"))
	t.Log(myAtoi("a+32"), myAtoi("a-32"))
	t.Log(myAtoi("+32ds"), myAtoi("-32d"))
	t.Log(myAtoi("12323490324302984032948"), myAtoi("-3249023483294932"))
	t.Log(myAtoi("9223372036854775808"))
	t.Log(myAtoi("-9223372036854775809"))
}

func TestPalindromeNumber(t *testing.T) {
	t.Log(isPalindromeNum(-121))
	t.Log(isPalindromeNum(121))
	t.Log(isPalindromeNum(10))
	t.Log(isPalindromeNum(1001))
	t.Log(isPalindromeNum(10001))
}

func TestCherryPickup(t *testing.T) {
	t.Log(cherryPickup([][]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, 1, 1},
	}))

	t.Log(cherryPickup([][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}))
	t.Log(cherryPickup([][]int{
		{1, 1, 1},
		{1, -1, 1},
		{1, 1, 1},
	}))

	t.Log(cherryPickup([][]int{
		{0, 1, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{-1, 1, 1, 1, -1},
		{0, 1, 1, 1, 0},
		{1, 0, -1, 0, 0},
	}))

	t.Log(cherryPickup([][]int{
		{1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1},
	}), 15)
}

func TestInsertionSort(t *testing.T) {
	N := 1000
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}

	insertionSort(arr)
	t.Log(arr)
}

func TestSelectionSort(t *testing.T) {
	N := 1000
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}

	SelectionSort(arr)
	t.Log(arr)
}

func TestBubbleSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	N := 1000
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}

	bubbleSort(arr)
	t.Log(arr)
}

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	N := 20
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}
	t.Log(arr)
	mergeSort(arr)
	t.Log(arr)

}

func TestQuickSoft3(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	N := 20
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}
	t.Log(arr)
	quickSort3(arr)
	t.Log(arr)
}

func TestShellSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	N := 20
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}
	t.Log(arr)
	shellSort(arr)
	t.Log(arr)

}

func TestCompareSort(t *testing.T) {
	N := 10000
	arrs := make([][]int, 3)
	for i := 0; i < 3; i++ {
		arrs[i] = make([]int, N)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		//v := rand.Intn(N*N)
		v := rand.Int()
		for _, arr := range arrs {
			arr[i] = v
		}
	}
	now := time.Now()
	QuickSoft(arrs[0])
	t.Log(arrs[0])
	t.Log("quick sort", time.Now().Sub(now))
	time.Sleep(time.Second)
	now = time.Now()
	quickSort3(arrs[1])
	t.Log(arrs[1])
	t.Log("quick sort3", time.Now().Sub(now))
	time.Sleep(time.Second)
	now = time.Now()
	shellSort(arrs[2])
	t.Log(arrs[2])
	t.Log("shell sort", time.Now().Sub(now))

}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(binarySearch(arr, 1))
	t.Log(binarySearch(arr, 3))
	t.Log(binarySearch(arr, 6))
	t.Log(binarySearch(arr, 8))
	t.Log(binarySearch(arr, 10))
	t.Log(binarySearch(arr, 11))
	t.Log(binarySearch(arr, 7))
}

func TestTopK(t *testing.T) {
	arr := []int{4, 2, 6, 4, 2, 6, 8, 9, 10, 3, 1, 3}
	t.Log(topk(arr, 1))
	t.Log(topk(arr, 2))
	t.Log(topk(arr, 3))
	t.Log(topk(arr, 4))
	t.Log(topk(arr, 5))
	t.Log(topk(arr, 6))
}
