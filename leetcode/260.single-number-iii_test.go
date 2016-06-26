package leetcode

import "testing"

func Test_singleNumber3(t *testing.T) {
	n := singleNumber3([]int{1, 2, 1, 3, 2, 5})
	if len(n) != 2 || n[0] != 3 || n[1] != 5 {
		t.Fatal(n)
	}
}
