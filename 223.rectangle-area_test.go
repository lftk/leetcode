package leetcode

import (
	"testing"
)

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	s := (C-A)*(D-B) + (G-E)*(H-F)
	if E >= C || A >= G || F >= D || H <= B {
		return s
	}
	x := (min(C, G) - max(A, E))
	y := (min(D, H) - max(B, F))
	return s - x*y
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func TestComputeArea(t *testing.T) {
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
