package leetcode

import (
	"testing"
)

func findKthLargest(nums []int, k int) int {
	num := nums[0]
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[r] <= num {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] >= num {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = num

	if l+1 == k {
		return num
	} else if l+1 > k {
		return findKthLargest(nums[:l+1], k)
	} else {
		return findKthLargest(nums[l+1:], k-l-1)
	}
}

func TestFindKthLargest(t *testing.T) {
	v := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	if v != 5 {
		t.Fatalf("%d != 5", v)
	}
}
