package leetcode

import "testing"

func Test_majorityElement(t *testing.T) {
	n := majorityElement([]int{1, 2, 3, 3, 1, 2, 1, 2, 1})
	if n != 1 {
		t.Fatalf("%d != 1", n)
	}
}
