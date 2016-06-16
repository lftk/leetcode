package leetcode

import "testing"

func Test_reverseWords(t *testing.T) {
	s1 := reverseWords("  the sky is blue  ")
	if s1 != "blue is sky the" {
		t.Fatal(s1)
	}
	s2 := reverseWords("a")
	if s2 != "a" {
		t.Fatal(s2)
	}
}
