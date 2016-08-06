package leetcode

import (
	"testing"
)

func Test_isPowerOfTwo(t *testing.T) {
	b1 := isPowerOfTwo(1024)
	if !b1 {
		t.Errorf("%d should is power of two", 1024)
	}
}
