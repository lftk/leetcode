package leetcode

import "testing"

func Test_computeArea(t *testing.T) {
	s := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)
	if s != 45 {
		t.Fatalf("%d != 45", s)
	}

	s = computeArea(-2, -2, 2, 2, -2, -2, 2, 2)
	if s != 16 {
		t.Fatalf("%d != 16", s)
	}

	s = computeArea(-2, -2, 2, 2, -1, -1, 1, 1)
	if s != 16 {
		t.Fatalf("%d != 16", s)
	}
}
