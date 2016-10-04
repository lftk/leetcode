package leetcode

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	n1 := reverse(-12304)
	if n1 != -40321 {
		t.Error(-12304, n1)
	}
}
