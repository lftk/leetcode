package leetcode

import (
	"testing"
)

func Test_reverseBits(t *testing.T) {
	n1 := reverseBits(43261596)
	if n1 != 964176192 {
		t.Error(n1, 964176192)
	}
}
