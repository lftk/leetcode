package leetcode

import "testing"

func Test_singleNumber(t *testing.T) {
	n := singleNumber([]int{1, 2, 3, 2, 1, 3, 4})
	if n != 4 {
		t.Fatalf("%d != 4", n)
	}
}
