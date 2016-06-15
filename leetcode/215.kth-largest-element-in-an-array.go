package leetcode

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
