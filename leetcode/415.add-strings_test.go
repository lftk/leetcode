package leetcode

import (
	"testing"
)

func Test_addStrings(t *testing.T) {
	s1 := addStrings("12345", "5")
	if s1 != "12350" {
		t.Error(s1, "12350")
	}
}
