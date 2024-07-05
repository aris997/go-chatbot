package matcher

import "math"

// func min2(a, b int) (m int) {
// 	if a <= b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

func min3(a, b, c int) (m int) {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	if c <= b && c <= a {
		return c
	}
	return math.MinInt
}
