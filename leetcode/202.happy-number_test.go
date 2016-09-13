package leetcode

import (
	"testing"
)

func Test_isHappy(t *testing.T) {
	b1 := isHappy(7)
	if b1 != true {
		t.Error(b1, 7)
	}

	b2 := isHappy(19)
	if b2 != true {
		t.Error(b2, 19)
	}
}
