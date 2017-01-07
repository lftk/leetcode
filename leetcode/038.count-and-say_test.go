package leetcode

import (
	"testing"
)

func Test_countAndSay(t *testing.T) {
	s4 := countAndSay(4)
	if s4 != "1211" {
		t.Error(s4, "1211")
	}

	s5 := countAndSay(5)
	if s5 != "111221" {
		t.Error(s5, "111221")
	}

	s8 := countAndSay(8)
	if s8 != "1113213211" {
		t.Error(s8, "1113213211")
	}
}
