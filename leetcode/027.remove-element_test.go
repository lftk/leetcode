package leetcode

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	n := removeElement(nums, 3)
	if n != 2 {
		t.Errorf("%d != 2", n)
		return
	}
	if nums[0] != 2 || nums[1] != 2 {
		t.Error(nums[0], nums[1])
	}
}
