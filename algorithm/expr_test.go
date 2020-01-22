package algorithm

import "testing"

func TestAlgoExpr(t *testing.T) {
	t.Log(evalAlgoExpr("1+2*3+5"))
	t.Log(evalAlgoExpr("(1+2*3+5)%5"))
	t.Log(evalAlgoExpr("-1*3+-2"))
	t.Log(evalAlgoExpr("(3+4)*(5-1)"))
}

func TestParseExpr(t *testing.T) {
	expr := []interface{}{
		"(",
		"(",
		true,
		"|",
		"!",
		false,
		")",
		"&",
		"(",
		false,
		"|",
		"!",
		true,
		")",
		")",
		"&",
		true,
	}
	rpnExpr, err := parse(expr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rpnExpr)

	t.Log(calc(rpnExpr))
}
