package algorithm

import (
	"fmt"
	"sort"
	"strconv"
)

/**
输入一个正整数的数组，把数组里面所有数字拼接成一个数，打印能拼接出的所有数字中最小的一个。
*/

func spliceArrayToMinNumber(arr []int) {
	if l := len(arr); l == 0 {
		fmt.Println(0)
	} else if l == 1 {
		fmt.Println(arr[0])
	}

	// 将数组元素按照一定的规则进行排序，比如元素n和m，排序后nm<mn
	sort.Slice(arr, func(i, j int) bool {
		si := strconv.Itoa(arr[i])
		sj := strconv.Itoa(arr[j])
		return si+sj < sj+si
	})

	// 直接拼接字符串或者打印，否则会溢出
	for _, n := range arr {
		fmt.Print(n)
	}
	fmt.Println()

}
