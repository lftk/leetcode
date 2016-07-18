package leetcode

import "testing"

func Test_getSum(t *testing.T) {
	s := getSum(1, 2)
	if s != 3 {
		t.Fatalf("%d != 3", s)
	}
}
