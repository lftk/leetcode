package leetcode

import "testing"

func Test_singleNumber2(t *testing.T) {
	n := singleNumber2([]int{1, 2, 1, 3, 1, 3, 3})
	if n != 2 {
		t.Fatalf("%d != 2", n)
	}
}
