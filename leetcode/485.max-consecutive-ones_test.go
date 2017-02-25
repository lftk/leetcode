package leetcode

import (
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	n1 := findMaxConsecutiveOnes([]int{1})
	if n1 != 1 {
		t.Error(n1, 1)
		return
	}

	n2 := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1})
	if n2 != 2 {
		t.Error(n2, 2)
		return
	}
}
