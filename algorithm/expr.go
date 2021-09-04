package algorithm

import (
	"errors"
	"fmt"
	"github.com/ymcvalu/ds_algo/data_structure/stack"
)

func parse(expr []interface{}) ([]interface{}, error) {
	rpnExpr := make([]interface{}, 0, len(expr))
	stack := stack.New()
	var last interface{}
	for _, it := range expr {
		switch it {
		case "(", "!":
			stack.Push(it)
		case ")":
			valid := false
			for !stack.IsEmpty() {
				v := stack.Pop()
				if v == "(" {
					valid = true
					break
				}
				rpnExpr = append(rpnExpr, v)
			}
			if !valid {
				return nil, errors.New("lack (")
			}
		case "&", "|":
			if last == "!" || last == "&" || last == "|" {
				return nil, errors.New(fmt.Sprintf("invalid expr: %s%s", last, it))
			}
			stack.Push(it)

		default:
			rpnExpr = append(rpnExpr, it)
			if last == "!" {
				rpnExpr = append(rpnExpr, stack.Pop())
			}
		}
		last = it
	}

	for !stack.IsEmpty() {
		top := stack.Pop()
		if top == "(" {
			return nil, errors.New("lack )")
		}
		rpnExpr = append(rpnExpr, top)
	}

	return rpnExpr, nil
}

func calc(rpnExpr []interface{}) bool {
	stack := stack.New()
	for _, it := range rpnExpr {
		switch it {
		case "&":
			v1 := stack.Pop()
			v2 := stack.Pop()
			stack.Push(v1 == true && v2 == true)
		case "|":
			v1 := stack.Pop()
			v2 := stack.Pop()
			stack.Push(v1 == true || v2 == true)

		case "!":
			v := stack.Pop()
			if v == true {
				stack.Push(false)
			} else {
				stack.Push(true)
			}
		default:
			stack.Push(it)
		}
	}
	return stack.Pop() == true
}
