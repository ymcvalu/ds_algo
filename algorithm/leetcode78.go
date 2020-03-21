package algorithm

// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）

func subsets(arr []int) [][]int {
	result := [][]int{}
	_subsets(arr, []int{}, &result)
	return result
}

func _subsets(arr []int, path []int, result *[][]int) {
	if len(arr) == 0 {
		*result = append(*result, append([]int(nil), path...))
		return
	}

	_subsets(arr[1:], path, result)
	_subsets(arr[1:], append(path, arr[0]), result)
}
