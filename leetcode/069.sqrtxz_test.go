package leetcode

import (
	"math"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n1 := mySqrt(i)
		n2 := int(math.Sqrt(float64(i)))
		if n1 != n2 {
			t.Fatal(i, n1, n2)
		}
	}
}
