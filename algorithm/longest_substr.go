package algorithm

/**
Given a string, find the length of the longest substring without repeating characters.

Example 1:
	Input: "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", which the length is 3.

Example 2:
	Input: "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

Example 3:
	Input: "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
		Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	rec := make([]int, len(s))
	rec[0] = 1
	max := 1
	for i := 1; i < len(s); i++ {
		lm := rec[i-1]
		j := i - 1
		for ; lm > 0; lm, j = lm-1, j-1 {
			if s[j] == s[i] {
				break
			}
		}
		if lm == 0 {
			rec[i] = rec[i-1] + 1

		} else {
			rec[i] = i - j
		}
		if max < rec[i] {
			max = rec[i]
		}

	}

	return max
}

