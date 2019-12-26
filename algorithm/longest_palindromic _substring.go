package algorithm

/**
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
Example 1:
	Input: "babad"
	Output: "bab"
	Note: "aba" is also a valid answer.

Example 2:
	Input: "cbbd"
	Output: "bb"
*/

func longestPalindrome1(s string) string {
	sl := len(s)
	if sl <= 1 {
		return s
	}

	var (
		mi, mj, max int
	)

	dp := make([][]int, sl)
	for i := range dp {
		dp[i] = make([]int, sl)
		dp[i][i] = 1
	}

	for d := 1; d < sl; d++ {
		for i := 0; i+d < sl; i++ {
			if s[i] == s[i+d] && dp[i+1][i+d-1] == d-1 {
				dp[i][i+d] = d + 1
				if d+2 > max {
					max = d + 1
					mi = i
					mj = i + d
				}
			}
		}
	}
	return s[mi : mj+1]
}

func longestPalindrome2(s string) string {
	sl := len(s)
	if sl <= 1 {
		return s
	}

	dp := make([][]int, sl)
	for i := range dp {
		dp[i] = make([]int, sl)
		dp[i][i] = 1
	}

	var maxI, maxJ, max int

	for l := 1; l < sl; l++ {
		for i := 0; i+l < sl; i++ {
			j := i + l
			if s[i] != s[j] {
				continue
			}
			if l == 1 {
				dp[i][j] = 2
			} else if dp[i+1][j-1] > 0 {
				dp[i][j] = dp[i+1][j-1] + 2
			}
			if dp[i][j] > max {
				max, maxI, maxJ = dp[i][j], i, j
			}
		}
	}
	return s[maxI : maxJ+1]
}
