package leetcode

import (
	"testing"
)

func Test_thirdMax(t *testing.T) {
	n1 := thirdMax([]int{1, 1, 2})
	if n1 != 2 {
		t.Error(n1)
	}
}
