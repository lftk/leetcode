package leetcode

import (
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	n1 := findContentChildren([]int{1, 2, 3}, []int{1, 1})
	if n1 != 1 {
		t.Error(n1, 1)
		return
	}

	n2 := findContentChildren([]int{1, 2}, []int{1, 2, 3})
	if n2 != 2 {
		t.Error(n2, 2)
		return
	}
}
