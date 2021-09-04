package algorithm

// minLexSubsequence 字符串的最小字典序子序列
// @str 指定字符串
// @n 子序列长度
func minLexSubsequence(str string, n int) string {
	if n >= len(str) {
		return str
	}

	strLen := len(str)
	seq := make([]byte, n)
	seqNextIdx := 0
	for idx := range str {
		c := str[idx]
		curIdx := seqNextIdx
		// 贪心策略：如果当前字符更小，并且后续子字符串长度足够，则替换已经选择的字符
		for i := seqNextIdx - 1; i >= 0; i-- {
			if c >= seq[i] || strLen-idx+i < n {
				break
			}
			curIdx = i
		}
		if curIdx < n {
			seq[curIdx] = c
			seqNextIdx = curIdx + 1
		}
	}

	return string(seq)
}
