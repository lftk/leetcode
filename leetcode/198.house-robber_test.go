package leetcode

import (
	"testing"
)

func Test_rob(t *testing.T) {
	n1 := rob([]int{3, 2, 1, 5})
	if n1 != 8 {
		t.Error(n1, 8)
	}
}
