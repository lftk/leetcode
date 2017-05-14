package leetcode

import (
	"testing"
)

func Test_twoSum2(t *testing.T) {
	r1 := twoSum2([]int{2, 3, 4}, 6)
	if len(r1) != 2 {
		t.Error(r1)
		return
	}
	if r1[0] != 1 || r1[1] != 3 {
		t.Error(r1)
		return
	}
}
