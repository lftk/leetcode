package leetcode

import (
	"testing"
)

func Test_reverseVowels(t *testing.T) {
	s1 := reverseVowels("hello")
	if s1 != "holle" {
		t.Error("hello", s1)
	}

	s2 := reverseVowels("leetcode")
	if s2 != "leotcede" {
		t.Error("leetcode", s2)
	}
}
