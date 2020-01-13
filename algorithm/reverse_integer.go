package algorithm

import "math"

/**
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321

Example 2:

Input: -123
Output: -321

Example 3:

Input: 120
Output: 21

Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/


func reverse(x int) int {
	if x == 0 || x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
	_x := x
	if x < 0 {
		_x = -x
	}

	py := 0
	y := 0
	for _x > 0 {
		py = y
		y = y*10 + (_x % 10)
		// overflow
		if y > math.MaxInt32 || (x < 0 && y > -math.MinInt32) || py*y < 0 {
			return 0
		}
		_x /= 10
	}

	if x < 0 {
		return -y
	}

	return y
}