package leetcode

import (
	"testing"
)

func Test_findNthDigit(t *testing.T) {
	n1 := findNthDigit(10)
	if n1 != 1 {
		t.Error(10, 1, n1)
	}
}
