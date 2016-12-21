package leetcode

import (
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	nums := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	for i, val := range []int{5, 6} {
		if val != nums[i] {
			t.Error(i, val, nums[i])
			break
		}
	}
}
