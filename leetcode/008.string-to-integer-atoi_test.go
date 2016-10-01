package leetcode

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	n1 := myAtoi("-1234")
	if n1 != -1234 {
		t.Error("-1234", n1)
	}
}
