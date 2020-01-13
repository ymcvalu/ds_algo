package algorithm

/**
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true

Example 2:

Input: "race a car"
Output: false
*/

func isPalindrome(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}
	var (
		i      = 0
		j      = len(s) - 1
		b1, b2 byte
	)

	for j > i {

		for b1 = lower(s[i]); i < j && b1 == ' '; b1 = lower(s[i]) {
			i++
		}
		for b2 = lower(s[j]); i < j && b2 == ' '; b2 = lower(s[j]) {
			j--
		}

		if i <= j && b1 != b2 {
			return false
		}
		i++
		j--
	}
	return true

}

func lower(b byte) byte {
	if b >= '0' && b <= '9' || b >= 'a' && b <= 'z' {
		return b
	} else if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return ' '
}
