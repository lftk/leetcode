package leetcode

import (
	"testing"
)

func Test_hammingDistance(t *testing.T) {
	n := hammingDistance(1, 4)
	if n != 2 {
		t.Error(n, 2)
	}
}
