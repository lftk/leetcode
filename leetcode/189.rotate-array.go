// https://leetcode.com/problems/rotate-array/

package leetcode

func rotate(nums []int, k int) {
	size := len(nums)
	if k %= size; k > 0 && k < size {
		rotateArray(nums[:size-k])
		rotateArray(nums[size-k:])
		rotateArray(nums)
	}
}

func rotateArray(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
