package algorithm

import "testing"

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
