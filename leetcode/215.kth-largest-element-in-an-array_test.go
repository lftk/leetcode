package leetcode

import "testing"

func Test_findKthLargest(t *testing.T) {
	v := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	if v != 5 {
		t.Fatalf("%d != 5", v)
	}
}
