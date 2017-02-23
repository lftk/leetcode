package leetcode

import (
	"testing"
)

func Test_convertToBase7(t *testing.T) {
	s1 := convertToBase7(100)
	if s1 != "202" {
		t.Error(100, s1, 202)
		return
	}

	s2 := convertToBase7(-7)
	if s2 != "-10" {
		t.Error(-7, s2, -10)
		return
	}
}
