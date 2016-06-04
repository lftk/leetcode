package leetcode

import (
	"math"
	"testing"
)

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

func TestMySqrt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n1 := mySqrt(i)
		n2 := int(math.Sqrt(float64(i)))
		if n1 != n2 {
			t.Fatal(i, n1, n2)
		}
	}
}
