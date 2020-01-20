package algorithm

func getNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	k := -1
	for i := 1; i < len(pattern); i++ {
		for k != -1 && pattern[k+1] != pattern[i] {
			// 当前字符不匹配，而前面[i-k-1,i]这k+1个字符刚好是模式串前缀
			// 这时候我们需要判断是否存在x>i-k-1，使得[x,i]刚好也是模式串前缀
			// 因为[i-k-1,i]是模式串的前缀，我们直接使用next数组
			// 如果(x=next[k])>=0，则[i-x-1,i]刚好也是模式串前缀
			k = next[k]
		}
		if pattern[k+1] == pattern[i] {
			k++
		}
		next[i] = k
	}
	return next
}

func kmp(str, pattern string) int {
	next := getNext(pattern)
	j := 0
	for i := 0; i < len(str); i++ {
		for j > 0 && str[i] != pattern[j] {
			j = next[j-1] + 1
		}
		if str[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			return i - len(pattern) + 1
		}
	}

	return -1
}
