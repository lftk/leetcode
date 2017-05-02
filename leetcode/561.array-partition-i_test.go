package leetcode

import (
	"testing"
)

func Test_arrayPairSum(t *testing.T) {
	s1 := arrayPairSum([]int{1, 4, 3, 2})
	if s1 != 4 {
		t.Error(s1, 4)
		return
	}
}
