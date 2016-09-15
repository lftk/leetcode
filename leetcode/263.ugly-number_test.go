package leetcode

import (
	"testing"
)

func Test_isUgly(t *testing.T) {
	b1 := isUgly(6)
	if !b1 {
		t.Error(b1, 6)
	}

	b2 := isUgly(8)
	if !b2 {
		t.Error(b2, 8)
	}

	b3 := isUgly(14)
	if b3 {
		t.Error(b3, 14)
	}
}
