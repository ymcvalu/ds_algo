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

func longestPalindrome(s string) string {
	sl := len(s)
	if sl <= 1 {
		return s
	}

	matrix := make([][]int, sl)
	for i := range matrix {
		matrix[i] = make([]int, sl)
	}

	var maxI, maxJ, max int

	for l := 0; l < sl; l++ {
		for i := 0; i+l < sl; i++ {
			j := i + l
			if s[i] != s[j] {
				continue
			}
			if j-i == 1 {
				matrix[i][j] = 2
			} else if j-i > 1 {
				if matrix[i+1][j-1] > 0 {
					matrix[i][j] = matrix[i+1][j-1] + 2
				}
			} else {
				matrix[i][j] = 1
			}
			if matrix[i][j] > max {
				max, maxI, maxJ = matrix[i][j], i, j
			}
		}
	}
	return s[maxI : maxJ+1]
}