package algorithm

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
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
	t.Log(isMatch("mississippi", "mis*is*P*.*"))
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
	quickSort3a(arr)
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

func TestLinkedListQuickSort(t *testing.T) {
	N := 20
	M := 100
	arr := make([]int, N)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(M)
	}
	head := arr2LinkedList(arr)

	LinkedListQuickSort(head)
	t.Log(arr)
	printList("soft", head)
}

func TestCompareSort(t *testing.T) {
	type SortAlgo struct {
		Func func([]int)
		Name string
	}

	algos := []SortAlgo{
		{
			Func: QuickSoft,
			Name: "quick sort",
		},
		{
			Func: quickSort3a,
			Name: "quick sort 3a",
		},
		{
			Func: quickSort3b,
			Name: "quick sort 3b",
		},
		{
			Func: shellSort,
			Name: "shell sort",
		},
	}

	N := 100000
	M := 100
	arrs := make([][]int, len(algos))
	for i := 0; i < len(algos); i++ {
		arrs[i] = make([]int, N)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		v := rand.Intn(M)
		for _, arr := range arrs {
			arr[i] = v
		}
	}

	for i, algo := range algos {
		time.Sleep(time.Second)
		now := time.Now()
		algo.Func(arrs[i])
		t.Log(algo.Name, time.Now().Sub(now))
	}

	if len(arrs) > 1 {
		for i := 1; i < len(arrs); i++ {
			if !reflect.DeepEqual(arrs[0], arrs[i]) {
				t.Errorf("%s is incorrent", algos[i].Name)
			}
		}
	}

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

func TestBM(t *testing.T) {
	t.Log(bm("abcdefg", "abc"))
	t.Log(bm("adjksd", "jks"))
	t.Log(bm("eidosljdls", "dls"))
}

func TestKMP(t *testing.T) {
	t.Log(kmp("abcdefg", "abc"))
	t.Log(kmp("adjksd", "jks"))
	t.Log(kmp("eidosljdls", "dls"))
}

func TestGetNext(t *testing.T) {
	t.Log(getNext("ababdababcaba"))
}

func TestFindNumInMatrix(t *testing.T) {
	t.Log(FindNumInMatrix([]int{1, 2, 8, 9, 2, 4, 9, 12, 4, 7, 10, 13, 6, 8, 11, 15}, 4, 4, 7))
	t.Log(FindNumInMatrix([]int{1, 2, 8, 9, 2, 4, 9, 12, 4, 7, 10, 13, 6, 8, 11, 15}, 4, 4, 20))
}

func TestTraverseBinaryTree(t *testing.T) {
	root := RecoverBinaryTree([]int{10, 6, 4, 8, 14, 12, 16}, []int{4, 6, 8, 10, 12, 14, 16})
	TraverseBinaryTree(root)
}

func TestStackQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Log(q.Pop())
	t.Log(q.Pop())
	q.Push(4)
	q.Push(5)
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
}

func TestQueueStack(t *testing.T) {
	stk := NewStack()
	stk.Push(1)
	stk.Push(2)
	t.Log(stk.Pop())
	stk.Push(3)
	t.Log(stk.Pop())
	t.Log(stk.Pop())
	t.Log(stk.Pop())
}

func TestFibonacci(t *testing.T) {
	t.Log(FibonacciA(30))
	t.Log(FibonacciB(30))
	t.Log(FibonacciC(30))
}

func TestMinOfRotateArr(t *testing.T) {
	t.Log(MinOfRotateArr([]int{3, 4, 5, 1, 2}))
	t.Log(MinOfRotateArr([]int{1, 2, 3, 4, 5}))
	t.Log(MinOfRotateArr([]int{5, 1, 2, 3, 4}))
}

func TestDumpWater(t *testing.T) {
	DumpWater()
}

func TestFindLadders(t *testing.T) {
	t.Log(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	t.Log(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}

func TestLCA(t *testing.T) {
	root := RecoverBinaryTree([]int{10, 6, 4, 8, 14, 12, 16}, []int{4, 6, 8, 10, 12, 14, 16})
	t.Log(LCA(root, 10, 6).Value, 10)
	t.Log(LCA(root, 4, 8).Value, 6)
	t.Log(LCA(root, 8, 12).Value, 10)
	t.Log(LCA(root, 4, 16).Value, 10)
	t.Log(LCA(root, 12, 16).Value, 14)
}

func TestTopologicalSort(t *testing.T) {
	g := buildGraph(true, 13, 0, 6, 0, 1, 0, 5, 2, 0, 2, 3, 3, 5, 5, 4, 6, 4, 6, 9, 7, 6, 8, 7, 9, 10, 9, 11, 9, 12, 11, 12)
	t.Log(topologicalSortA(g))
	t.Log(topologicalSortB(g))
}

func TestKruskal(t *testing.T) {
	t.Log(kruskal(8, []Edge{
		{0, 7, 0.16},
		{2, 3, 0.17},
		{1, 7, 0.19},
		{0, 2, 0.26},
		{5, 7, 0.28},
		{1, 3, 0.29},
		{1, 5, 0.32},
		{2, 7, 0.34},
		{4, 5, 0.35},
		{1, 2, 0.36},
		{4, 7, 0.37},
		{0, 4, 0.38},
		{6, 2, 0.40},
		{2, 6, 0.52},
		{6, 0, 0.58},
		{6, 4, 0.93},
	}))
}

func TestDijkstraSP(t *testing.T) {
	g := buildWeightGraph(true, 8, []Edge{
		{0, 1, 5},
		{0, 3, 2},
		{0, 5, 1},
		{2, 1, 1},
		{3, 2, 2},
		{4, 1, 1},
		{5, 4, 2},
		{5, 6, 3},
		{5, 7, 5},
		{6, 7, 1},
	})
	dist := dijkstraSP(g, 0)
	t.Log(dist)
}

func TestBellmanFordSP(t *testing.T) {
	g := buildWeightGraph(true, 8, []Edge{
		{0, 1, 5},
		{0, 3, 2},
		{0, 5, 1},
		{2, 1, 1},
		{3, 2, 2},
		{4, 1, 1},
		{5, 4, 2},
		{5, 6, 3},
		{5, 7, -5},
		{6, 7, 1},
	})
	dist := BellmanFord(g, 0)
	t.Log(dist)

	g = buildWeightGraph(true, 3, []Edge{
		{0, 1, 1},
		{0, 2, 2},
		{2, 1, -2},
	})
	t.Log(BellmanFord(g, 0))
}

func TestFloydWarshall(t *testing.T) {
	g := buildWeightGraph(true, 8, []Edge{
		{0, 1, 5},
		{0, 3, 2},
		{0, 5, 1},
		{2, 1, 1},
		{3, 2, 2},
		{4, 1, 1},
		{5, 4, 2},
		{5, 6, 3},
		{5, 7, -5},
		{6, 7, 1},
	})
	dp := FloydWarshall(g)
	t.Log(dp[0])

	g = buildWeightGraph(true, 4, []Edge{
		{0, 2, -2},
		{1, 0, 4},
		{1, 2, 3},
		{3, 1, -1},
		{2, 3, 2},
	})
	t.Log(FloydWarshall(g))

	g = buildWeightGraph(true, 5, []Edge{
		{0, 1, 1},
		{0, 2, 4},
		{0, 3, 1},
		{1, 2, 2},
		{2, 4, 3},
		{3, 2, 1},
	})
	t.Log(FloydWarshall(g)[0])

	g = buildWeightGraph(true, 3, []Edge{
		{0, 1, 1},
		{1, 2, -1},
		{2, 0, -1},
	})
	t.Log(FloydWarshall(g))
}

func TestFisherYatesShuffle(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 10)
		t.Log(FisherYatesShuffle(10))
	}
}

func TestSplitRope(t *testing.T) {
	t.Log(splitRopeDP(2))
	t.Log(splitRopeDP(3))
	t.Log(splitRopeDP(6))
	t.Log(splitRopeDP(7))
	t.Log(splitRopeDP(8))

}

func TestMorrisTraverseTreeInOrder(t *testing.T) {
	root := RecoverBinaryTree([]int{10, 6, 4, 8, 14, 12, 16}, []int{4, 6, 8, 10, 12, 14, 16})
	morrisTraverseInOrder(root)
	fmt.Println()
	InOrder(root)
	fmt.Println()
}

func TestMorrisTraverseTreePreOrder(t *testing.T) {
	root := RecoverBinaryTree([]int{10, 6, 4, 8, 14, 12, 16}, []int{4, 6, 8, 10, 12, 14, 16})
	morrisTraversePreOrder(root)
	fmt.Println()
	PreOrder(root)
	fmt.Println()
}

func TestMorrisTraverseTreePostOrder(t *testing.T) {
	root := RecoverBinaryTree([]int{10, 6, 4, 8, 14, 12, 16}, []int{4, 6, 8, 10, 12, 14, 16})
	morrisTraversePostOrder(root)
	fmt.Println()
	PostOrder(root)
	fmt.Println()
}

func TestFindGreatestSumOfSubarray(t *testing.T) {
	t.Log(findGreatestSumOfSubarray([]int{1, -2, 3, 10, -4, 7, 2, -5}))
}

func TestSpliceArrayToMinNumber(t *testing.T) {
	spliceArrayToMinNumber([]int{3, 32, 321})
}

func TestInversePairs(t *testing.T) {
	arrs := [][]int{
		{1, 2, 3, 4, 5, 6}, //0
		{7, 5, 6, 4},       // 5
		{7, 5, 6, 4, 1},    // 9
		{7, 2, 4, 6, 4},    // 5
		{7, 2, 5, 6, 4, 1}, // 11
	}

	for _, arr := range arrs {
		t.Log(InversePairs(arr), arr)
	}
}

func TestTreeSerialize(t *testing.T) {
	root := RecoverBinaryTree([]int{10, 6, 4, 8, 14, 12, 16}, []int{4, 6, 8, 10, 12, 14, 16})
	s := treeSerialize(root)
	t.Log(s)

	newTree := treeDeserialize(s)

	PreOrder(newTree)
	InOrder(newTree)
}

func TestDreamOfMondriann(t *testing.T) {
	type Case struct {
		N int
		M int
		R int
	}

	cases := []Case{
		{1, 2, 1},
		{1, 3, 0},
		{1, 4, 1},
		{2, 2, 2},
		{2, 3, 3},
		{2, 4, 5},
		{2, 11, 144},
		{4, 11, 51205},
	}
	for _, c := range cases {
		r := dreamOfMondriann(c.N, c.M)
		t.Log(c, r, c.R == r)
	}
}

func TestTreeZigZag(t *testing.T) {
	root := RecoverBinaryTree([]int{10, 7, 6, 8, 15, 13, 16}, []int{6, 7, 8, 10, 13, 15, 16})
	TreeZigZag(root)

	root = RecoverBinaryTree([]int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}, []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15})
	TreeZigZag(root)
}

func TestStringMinWindow(t *testing.T) {
	cases := []struct {
		s string
		t string
		r string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"ABCDEFGFJAKLD", "CE", "CDE"},
		{"ABCDEDFSDSDSD", "BDE", "BCDE"},
		{"aa", "aa", "aa"},
	}
	for _, c := range cases {
		if r := stringMinWindow(c.s, c.t); c.r != r {
			t.Errorf("failed pass %s %s, %s != %s", c.s, c.t, r, c.r)
		}
	}
}

func TestFindSubstring(t *testing.T) {
	cases := []struct {
		s     string
		words []string
		r     []int
	}{
		{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			r:     []int{0, 9},
		}, {
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			r:     []int{},
		}, {
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			r:     []int{},
		},
		{
			s:     "wordgoodgoodgoodbestwordword",
			words: []string{"word", "good", "best", "word"},
			r:     []int{12},
		},
		{
			s:     "barfoofoobarthefoobarman",
			words: []string{"bar", "foo", "the"},
			r:     []int{6, 9, 12},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "good"},
			r:     []int{8},
		},
	}
	for _, c := range cases {
		if r := findSubstring(c.s, c.words); !reflect.DeepEqual(r, c.r) {
			t.Errorf("failed to pass %s %v, %v != %v", c.s, c.words, r, c.r)
		}
	}
}

func TestMinSubArrayLen(t *testing.T) {
	cases := []struct {
		s    int
		nums []int
		r    int
	}{
		{
			s:    7,
			nums: []int{2, 3, 1, 2, 4, 3},
			r:    2,
		},
		{
			s:    6,
			nums: []int{10, 2, 3},
			r:    1,
		},
	}
	for _, c := range cases {
		if r := minSubArrayLen(c.s, c.nums); r != c.r {
			t.Errorf("failed to pass %d %v, %d != %d", c.s, c.nums, c.r, r)
		}
	}
}

func TestFindLength(t *testing.T) {
	cases := []struct {
		A []int
		B []int
		r int
	}{
		{
			A: []int{1, 2, 3, 2, 1},
			B: []int{3, 2, 1, 4, 7},
			r: 3,
		},
	}

	for _, c := range cases {
		if r := findLength(c.A, c.B); r != c.r {
			t.Errorf("faield to pass %v %v, %d != %d", c.A, c.B, c.r, r)
		}
	}
}

func TestMinStack(t *testing.T) {
	minStk := ConstructorMiniStack()
	minStk.Push(-2)
	minStk.Push(0)
	minStk.Push(-3)
	t.Log(minStk.GetMin())
	minStk.Pop()
	t.Log(minStk.GetMin())
	t.Log(minStk.Top())
	minStk.Pop()
	t.Log(minStk.GetMin())
}
