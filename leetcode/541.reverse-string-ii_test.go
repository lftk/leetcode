package leetcode

import (
	"testing"
)

func Test_reverseStr(t *testing.T) {
	s2 := reverseStr("abcdefg", 2)
	if s2 != "bacdfeg" {
		t.Error("abcdefg", s2, "bacdfeg")
	}
}
