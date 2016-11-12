package leetcode

import (
	"testing"
)

func Test_minMoves(t *testing.T) {
	n1 := minMoves([]int{1, 2, 3})
	if n1 != 3 {
		t.Error(n1, 3)
	}
}
