package leetcode

import (
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	n1 := hammingWeight(11)
	if n1 != 3 {
		t.Error(n1, 11)
	}
}
