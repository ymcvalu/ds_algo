package algorithm

// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

func permute(arr []int) [][]int {
	result := [][]int{}
	_permute(arr, 0, &result)
	return result
}

func _permute(arr []int, i int, result *[][]int) {
	if i == len(arr) {
		*result = append(*result, append([]int(nil), arr...))
		return
	}

	for j := i; j < len(arr); j++ {
		arr[j], arr[i] = arr[i], arr[j]
		_permute(arr, i+1, result)
		arr[j], arr[i] = arr[i], arr[j]
	}
}
