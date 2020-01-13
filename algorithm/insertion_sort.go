package algorithm

func insertionSort(arr []int) {
	ln := len(arr)
	if ln < 2 {
		return
	}

	for i := 1; i < ln; i++ {
		v := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > v; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = v
	}
}

// 排序的本质就是消除逆序对
// n个元素的随机数组，逆序对是O(n^2)
// 如果采用交换相连元素，一趟循环只能消除O(n)个逆序对，因此时间复杂度为O(n^2)
// 而如果想要突破O(n^2)，则需要一次交换或者一趟循环能够消除多个逆序对
// 比如快排，一次partition，将数组分为前后两部分，后面部分的元素大于前面部分元素，因此可以消除O(N/2*N/2)的有序对
// shell排序通过跨间隔比较，能够一次交换消除多个有序对
func shellSort(arr []int) {
	ln := len(arr)
	if ln < 2 {
		return
	}

	for g := 11; g > 0; g -= 2 {
		for i := g; i < ln; i++ {
			v := arr[i]
			j := i - g
			for ; j >= 0 && arr[j] > v; j -= g {
				arr[j+g] = arr[j]
			}
			arr[j+g] = v
		}
	}

}
