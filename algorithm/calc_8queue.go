package algorithm

import "fmt"

func queue_8() {
	result := make([]int, 8)
	for col := 0; col < 8; col++ {
		result[0] = col
		calc8Queue(1, result)
	}
}

func calc8Queue(row int, result []int) {
	if row == 8 {
		fmt.Println(result)
		return
	}

loop:
	for col := 0; col < 8; col++ {
		for i := 0; i < row; i++ {
			if result[i] == col || col-result[i] == row-i || col-result[i] == i-row {
				continue loop
			}
		}
		result[row] = col
		calc8Queue(row+1, result)
	}
}
