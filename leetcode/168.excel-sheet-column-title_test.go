package leetcode

import "testing"

func Test_convertToTitle(t *testing.T) {
	s1 := convertToTitle(28)
	if s1 != "AB" {
		t.Fatalf("%s != AB", s1)
	}
	s2 := convertToTitle(26)
	if s2 != "Z" {
		t.Fatalf("%s != Z", s2)
	}
}
