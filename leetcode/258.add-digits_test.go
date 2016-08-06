package leetcode

import (
	"testing"
)

func Test_addDigits(t *testing.T) {
	n1 := addDigits(1111)
	if n1 != 4 {
		t.Errorf("%d != 4", n1)
	}
}
