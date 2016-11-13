package leetcode

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	n1 := longestPalindrome("abccccdd")
	if n1 != 7 {
		t.Error(n1, 7)
	}
}
