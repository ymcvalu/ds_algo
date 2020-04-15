package algorithm

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	m := [256]int{}
	max := 1
	m[s[0]] = 1
	l, r := 0, 0
	for {
		r++
		if r == len(s) {
			break
		}

		if m[s[r]] > 0 && m[s[r]] > l {
			if r-l > max {
				max = r - l
			}
			l = m[s[r]]
		}

		m[s[r]] = r + 1
	}
	if r-l > max {
		max = r - l
	}
	return max
}
