package leetcode

import (
	"testing"
)

func Test_isAnagram(t *testing.T) {
	b1 := isAnagram("anagram", "nagaram")
	if !b1 {
		t.Error("anagram <-> nagaram")
	}
}
