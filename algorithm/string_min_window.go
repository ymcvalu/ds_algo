package algorithm

/**
76.最小覆盖子串
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：
输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"

说明：
如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/
func stringMinWindow(s string, t string) string {
	hash := [256]int{} // byte范围0-255
	for i := range t {
		hash[t[i]] += 1
	}

	min := ""
	l := 0
	cnt := 0
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			cnt++
		}
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}

		if cnt == len(t) && min == "" || len(min) > r-l+1 {
			min = s[l : r+1]
		}

	}

	return min
}
