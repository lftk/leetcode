package leetcode

import (
	"testing"
)

func Test_reverseWords3(t *testing.T) {
	s1 := reverseString("Let's take LeetCode contest")
	if s1 != "s'teL ekat edoCteeL tsetnoc" {
		t.Error(s1, "s'teL ekat edoCteeL tsetnoc")
		return
	}
}
