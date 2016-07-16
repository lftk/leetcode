package leetcode

import "testing"

func Test_missingNumber(t *testing.T) {
	n := missingNumber([]int{0, 1, 3})
	if n != 2 {
		t.Fatalf("%d != 2", n)
	}
}
