package leetcode

import "testing"

func Test_uniquePaths(t *testing.T) {
	n1 := uniquePaths(3, 7)
	if n1 != 28 {
		t.Fatalf("%d != 28", n1)
	}
	n2 := uniquePaths(51, 9)
	if n2 != 1916797311 {
		t.Fatalf("%d != 1916797311", n2)
	}
}
