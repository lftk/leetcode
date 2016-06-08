package leetcode

import "testing"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if j := nums[i] - 1; j >= 0 && j < n {
			if nums[i] != nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				i--
			}
		}
	}
	for i := 0; i < n; i++ {
		if i != nums[i]-1 {
			return i + 1
		}
	}
	return n + 1
}

func TestFirstMissingPositive(t *testing.T) {
	if n := firstMissingPositive([]int{1, 2, 0}); n != 3 {
		t.Fatalf("%d != 3", n)
	}
	if n := firstMissingPositive([]int{3, 4, -1, 1}); n != 2 {
		t.Fatalf("%d != 2", n)
	}
	if n := firstMissingPositive([]int{1, 2, 3}); n != 4 {
		t.Fatalf("%d != 4", n)
	}
}
