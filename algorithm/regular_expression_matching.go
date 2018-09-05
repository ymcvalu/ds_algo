package algorithm

/**
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:
 - s could be empty and contains only lowercase letters a-z.
 - p could be empty and contains only lowercase letters a-z, and characters like . or *.

Example 1:
Input:
	s = "aa"
	p = "a"
Output: false
	Explanation: "a" does not match the entire string "aa".

Example 2:
Input:
	s = "aa"
	p = "a*"
Output: true
	Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

 */

func isMatch(s string, p string) bool {
	sl, pl := len(s), len(p)
	si, pi := 0, 0
	for pi < pl {
		if si < sl && (s[si] == p[pi] || p[pi] == '.') {
			if pi+1 < pl && p[pi+1] == '*' {
				return isMatch(s[si+1:], p[pi:]) || isMatch(s[si:], p[pi+2:])
			}
			si++
			pi++
		} else if pi+1 < len(p) && p[pi+1] == '*' {
			pi += 2
		} else {
			return false
		}

	}
	return si == sl
}
