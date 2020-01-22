package algorithm

import (
	"leetcode/data_structure/stack"
	"strconv"
)

func evalAlgoExpr(str string) int {
	priM := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 3,
		'/': 3,
		'%': 3,
	}

	opStk := stack.New()
	numStk := stack.New()
	status := 1
	neg := false
	for i := 0; i < len(str); {
		if !isNum(str[i]) {
			op := str[i]
			if status == 2 || status == 1 && op != '-' && op != '+' && op != '(' {
				panic("invalid expr")
			}
			if status == 1 && op == '(' {
				status = 0
			}

			if status == 1 {
				status = 2
				if str[i] == '-' {
					neg = true
				}
				i++
				continue
			}

			if op != ')' {
				status = 1
			}

			for {
				if op == '(' || opStk.IsEmpty() {
					opStk.Push(op)
					break
				}
				pOp := opStk.Peek().(byte)
				if pOp == '(' && op == ')' {
					opStk.Pop()
					break
				}
				if priM[pOp] < priM[op] {
					opStk.Push(op)
					break
				}
				opStk.Pop()

				numStk.Push(evalOp(pOp, numStk.Pop().(int), numStk.Pop().(int)))
			}

			i++
			continue
		}
		status = 0
		j := i + 1
		for j < len(str) && isNum(str[j]) {
			j++
		}
		num, _ := strconv.Atoi(str[i:j])
		if neg {
			num *= -1
			neg = false
		}
		i = j
		numStk.Push(num)
	}

	for !opStk.IsEmpty() {
		numStk.Push(evalOp(opStk.Pop().(byte), numStk.Pop().(int), numStk.Pop().(int)))
	}

	if numStk.IsEmpty() {
		return 0
	}
	return numStk.Pop().(int)
}

func evalOp(op byte, op2, op1 int) int {
	switch op {
	case '+':
		return op1 + op2
	case '-':
		return op1 - op2
	case '*':
		return op1 * op2
	case '/':
		return op1 / op2
	case '%':
		return op1 / op2
	}
	panic("invalid op")
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}
