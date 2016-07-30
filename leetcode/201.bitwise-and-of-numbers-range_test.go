package leetcode

import (
	"testing"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	n := rangeBitwiseAnd(5, 7)
	if n != 4 {
		t.Errorf("%d != 4", n)
	}
}
