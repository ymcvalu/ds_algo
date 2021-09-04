package algorithm

import "testing"

func TestMinLexSubquence(t *testing.T) {
	cases := []struct {
		str string
		n   int
		seq string
	}{
		{
			str: "ababab",
			n:   3,
			seq: "aaa",
		},
		{
			str: "bbbqqaa",
			n:   2,
			seq: "aa",
		},
		{
			str: "bbbqqaa",
			n:   4,
			seq: "bbaa",
		},
		{
			str: "aaabcaca",
			n:   6,
			seq: "aaaaca",
		},
		{
			str: "gxebrgcaxy",
			n:   4,
			seq: "baxy",
		},
	}

	for _, c := range cases {
		if seq := minLexSubsequence(c.str, c.n); seq != c.seq {
			t.Errorf("case not pass: %+v, result is %s", c, seq)
		}
	}
}
