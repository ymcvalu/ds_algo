package algorithm

import (
	"math"
)

/**
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true

Example 2:

Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Follow up:

Coud you solve it without converting the integer to a string?
*/

func isPalindromeNum(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	i := int(math.Pow10(int(math.Log10(float64(x)))))

	j := 1

	for i >= j {
		if x/i%10 == x/j%10 {
			i /= 10
			j *= 10
			continue
		}
		return false
	}

	return true
}
