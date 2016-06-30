package leetcode

import "testing"

func Test_titleToNumber(t *testing.T) {
	n1 := titleToNumber("AA")
	if n1 != 27 {
		t.Fatalf("%d != 27", n1)
	}
	n2 := titleToNumber("A")
	if n2 != 1 {
		t.Fatalf("%d != 1", n2)
	}
}
