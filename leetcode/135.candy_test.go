package leetcode

import "testing"

func Test_candy(t *testing.T) {
	n := candy([]int{1, 2, 3, 2, 1, 4, 5, 2, 1, 4})
	if n != 19 {
		t.Fatalf("%d != 19", n)
	}
}
