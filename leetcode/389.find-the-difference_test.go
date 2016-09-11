package leetcode

import (
	"testing"
)

func Test_findTheDifference(t *testing.T) {
	c1 := findTheDifference("abcd", "abcde")
	if c1 != 'e' {
		t.Error(c1, 'e')
	}
}
