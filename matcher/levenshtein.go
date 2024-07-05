package matcher

func Distance(s1 string, s2 string) (value int) {
	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}
	if s1[0] == s2[0] {
		return Distance(s1[1:], s2[1:])
	}
	eval1 := Distance(s1[1:], s2)
	eval2 := Distance(s1, s2[1:])
	eval3 := Distance(s1[1:], s2[1:])

	return 1 + min3(eval1, eval2, eval3)
}
