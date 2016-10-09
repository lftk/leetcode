package leetcode

import (
	"testing"
)

func Test_addBinary(t *testing.T) {
	s1 := addBinary("11", "1")
	if s1 != "100" {
		t.Error(s1, 100)
	}
}
