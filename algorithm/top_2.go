package algorithm

// 求数组中第二大的数
// 使用插入排序，先找最大的，然后找第二大的，时间复杂度为O(n)

func Top2(arr []int) int {
	if len(arr) < 2 {
		panic("invalid arr")
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[1]
		} else {
			return arr[0]
		}
	}

	for i := 0; i < 2; i++ {
		v := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > v {
				v, arr[j] = arr[j], v
			}
		}
		arr[i] = v
	}
	return arr[1]
}
