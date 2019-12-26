package algorithm

/**
字符串S由小写字母构成，长度为n。定义一种操作，每次都可以挑选字符串中任意的两个相邻字母进行交换。询问在至多交换m次之后，字符串中最多有多少个连续的位置上的字母相同？
输入描述:
第一行为一个字符串S与一个非负整数m。(1 <= |S| <= 1000, 1 <= m <= 1000000)
输出描述:
一个非负整数，表示操作之后，连续最长的相同字母数量。
*/

// 解题思路：区间DP
// 首先获取每个字符的下标数组
// 将字符往中间靠拢，所需的交换次数最低
// dp[i][j]表示将第i个到第j个之间的字符连到一起需要的交换次数
// 则dp[i][j]=dp[i+1][j-1] + [idx(j)-idx(i)-1- len]

func swapLetter(str string, m int) int {
	letters := make(map[byte][]int, 26)
	for i := range str {
		letters[str[i]] = append(letters[str[i]], i)
	}

	var max int
	for letter := range letters {
		idxs := letters[letter]
		ln := len(idxs)

		if ln == 1 {
			if ln > max {
				max = ln
			}
			continue
		}
		dp := make([][]int, ln)
		for i := 0; i < ln; i++ {
			dp[i] = make([]int, ln)
		}

		for pl := 1; pl < ln; pl++ {
			for i := 0; i < ln-pl; i++ {
				dp[i][i+pl] = dp[i+1][i+pl-1] + idxs[i+pl] - idxs[i] - pl
				if dp[i][i+pl] <= m && pl+1 > max {
					max = pl + 1
				}
			}
		}

	}
	return max
}
