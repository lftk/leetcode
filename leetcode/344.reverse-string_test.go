package leetcode

import "testing"

func Test_reverseString(t *testing.T) {
	s := reverseString("hello")
	if s != "olleh" {
		t.Fatalf("%s != olleh", s)
	}
}
