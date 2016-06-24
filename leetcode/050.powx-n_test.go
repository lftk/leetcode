package leetcode

import (
	"math"
	"testing"
)

func Test_myPow(t *testing.T) {
	for x := -100.0; x < 100.0; x += 2.3 {
		for n := 1; n < 6; n++ {
			p1 := myPow(x, n)
			p2 := math.Pow(x, float64(n))
			if math.Abs(p1-p2) > 0.0001 {
				t.Fatal(p1, p2)
			}
		}
	}
}
