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

	s3 := reverseVowels("a.b,.")
	if s3 != "a.b,." {
		t.Error("a.b,.", s3)
	}
}
