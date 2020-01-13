package algorithm

/**
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
 */

func longestValidParentheses(s string) int {
	max := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if i-dp[i-1]-1 >= 0 && s[i-1-dp[i-1]] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 > 0 {
					dp[i] += dp[i-2-dp[i-1]]
				}
			}

			if dp[i] > max {
				max = dp[i]
			}

		}
	}
	return max
}
