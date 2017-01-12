package leetcode

import (
	"testing"
)

func Test_findComplement(t *testing.T) {
	n1 := findComplement(5)
	if n1 != 2 {
		t.Error(n1, 2)
	}

	n2 := findComplement(24)
	if n2 != 7 {
		t.Error(n2, 7)
	}
}
