package algorithm

import (
	"bytes"
)

/**
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:
string convert(string s, int numRows);

Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
 */

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]byte, numRows)
	idx := 0
	dir := 1
	for i := 0; i < len(s); i++ {
		rows[idx] = append(rows[idx], s[i])
		if dir == 1 && idx == numRows-1 {
			dir = -1
		} else if dir == -1 && idx == 0 {
			dir = 1
		}
		idx += dir
	}
	buf := bytes.NewBuffer(nil)
	for i := 0; i < numRows; i++ {
		buf.Write(rows[i])
	}
	return buf.String()
}
