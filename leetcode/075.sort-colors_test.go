package leetcode

import (
	"testing"
)

func Test_sortColors(t *testing.T) {
	c1 := []int{2, 1, 0}
	sortColors(c1)
	for i, c := range c1 {
		if c != i {
			t.Error(i, c)
		}
	}
}
