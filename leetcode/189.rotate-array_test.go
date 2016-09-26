package leetcode

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	for i, n := range []int{5, 6, 7, 1, 2, 3, 4} {
		if n != nums[i] {
			t.Error(i, n, nums[i])
			break
		}
	}
}
