package leetcode

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, r := 0, x
	for l < r {
		m := (r-l)/2 + l
		if m == l || m == r {
			return m
		}
		mm := m * m
		if mm == x {
			return m
		}
		if mm > x {
			r = m
		} else {
			l = m
		}
	}
	return l
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
