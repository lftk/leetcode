package leetcode

import (
	"testing"
)

func Test_convert(t *testing.T) {
	s1 := convert("PAYPALISHIRING", 3)
	if s1 != "PAHNAPLSIIGYIR" {
		t.Error(s1, "PAHNAPLSIIGYIR")
	}
}
