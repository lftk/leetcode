package leetcode

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	n := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	if n != 2 {
		t.Fatalf("%d != 2", n)
	}
}
