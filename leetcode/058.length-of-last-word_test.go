package leetcode

import (
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	n1 := lengthOfLastWord("Hello World  ")
	if n1 != 5 {
		t.Error(n1, "Hello World  ")
	}
}
