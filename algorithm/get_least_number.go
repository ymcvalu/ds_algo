package algorithm

/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
输出数据不要求排序。

解题思路
方法一：先升序排序一下，然后直接选择最小的k个数，题目没要求结果要有序，没必要全排序，使用快排的话时间复杂度是O(n log n)
方法二：使用最小堆，先构造堆，然后连续取k次，堆的构造时间复杂度是O(n)，堆的维护是O(log n)，总的时间复杂度是O(n + klog n)
方法三：使用partition函数，因为输出数据不需要排序
*/

func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 {
		return nil
	}
	ln := len(arr)
	if ln < k {
		return nil
	} else if ln == k {
		return arr
	}

	m := getLeastNumbersPartition(arr)
	if m == k {
		return arr[:m]
	} else if m+1 == k {
		return arr[:m+1]
	} else if k > m+1 {
		return append(arr[:m+1], getLeastNumbers(arr[m+1:], k-m-1)...)
	} else {
		return getLeastNumbers(arr[:m], k)
	}

}

func getLeastNumbersPartition(arr []int) int {
	ln := len(arr)
	if ln == 0 {
		return -1
	}
	if ln == 1 {
		return 0
	}

	v := arr[0]

	p, q := 1, ln-1
	for p < q {

		for p < ln && arr[p] <= v {
			p++
		}

		for q > 0 && arr[q] >= v {
			q--
		}

		if p < q {
			arr[p], arr[q] = arr[q], arr[p]
		}
	}

	arr[q], arr[0] = arr[0], arr[q]

	return q
}
