package algorithm

func strIncludeSeq(str string, seq string) bool {
	p, q := 0, 0
	for p < len(seq) && q < len(str) {
		if seq[p] == str[q] {
			p++
		}
		q++
	}
	if p == len(seq) {
		return true
	}
	return false
}
