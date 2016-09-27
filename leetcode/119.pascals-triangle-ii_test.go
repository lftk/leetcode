package leetcode

import (
	"testing"
)

func Test_getRow(t *testing.T) {
	n3 := getRow(3)
	for i, n := range []int{1, 3, 3, 1} {
		if i >= len(n3) || n != n3[i] {
			t.Error(i, n, n3[i])
			break
		}
	}
}
