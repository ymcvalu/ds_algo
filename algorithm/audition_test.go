package algorithm

import "testing"

func TestVerify(t *testing.T) {
	t.Log(verify("people"))
	t.Log(verify("peopple"))
	t.Log(verify("ppeoplplle"))
	t.Log(verify("peopl"))
}
