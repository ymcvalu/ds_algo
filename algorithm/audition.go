package algorithm

func verify(str string) bool {
	people := "people"
	p, q := 0, 0
	for p < len(people) && q < len(str) {
		if people[p] == str[q] {
			p++
		}
		q++
	}
	if p == len(people) {
		return true
	}
	return false
}
