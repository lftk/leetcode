package leetcode

import "testing"

func Test_merge2(t *testing.T) {
	nums1 := make([]int, 5)
	nums1[0] = 3
	nums1[1] = 4
	nums1[2] = 5
	nums2 := []int{1, 2}
	merge2(nums1, 3, nums2, 2)
	if nums1[0] != 1 || nums1[4] != 5 {
		t.Fatal(nums1)
	}
}
