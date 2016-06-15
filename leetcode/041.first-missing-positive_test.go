package leetcode

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	if n := firstMissingPositive([]int{1, 2, 0}); n != 3 {
		t.Fatalf("%d != 3", n)
	}
	if n := firstMissingPositive([]int{3, 4, -1, 1}); n != 2 {
		t.Fatalf("%d != 2", n)
	}
	if n := firstMissingPositive([]int{1, 2, 3}); n != 4 {
		t.Fatalf("%d != 4", n)
	}
}
