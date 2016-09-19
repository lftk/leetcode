package leetcode

import (
	"testing"
)

func Test_plusOne(t *testing.T) {
	d1 := plusOne([]int{9, 9})
	for i, n := range []int{1, 0, 0} {
		if i >= len(d1) || n != d1[i] {
			t.Error(i, n, d1[i])
			break
		}
	}
}
