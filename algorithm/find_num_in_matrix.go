package algorithm

/**
	matrix:每行递增，每列递增
	1  2  8  9
	2  4  9  12
    4  7  10 13
    6  8  11 15
*/
func FindNumInMatrix(matrix []int, rows, cols int, num int) bool {
	r := 0
	c := cols - 1
	for r < rows && c >= 0 {
		if n := matrix[r*cols+c]; n == num {
			return true
		} else if n > num {
			c--
		} else {
			r++
		}
	}

	return false
}

