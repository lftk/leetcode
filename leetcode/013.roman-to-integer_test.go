package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	n := romanToInt("DCXXI")
	if n != 621 {
		t.Fatalf("%d != 621", n)
	}
}
