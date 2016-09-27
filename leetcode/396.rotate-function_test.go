package leetcode

import (
	"testing"
)

func Test_maxRotateFunction(t *testing.T) {
	n1 := maxRotateFunction([]int{4, 3, 2, 6})
	if n1 != 26 {
		t.Error([]int{4, 3, 2, 6}, n1, 26)
	}
}
