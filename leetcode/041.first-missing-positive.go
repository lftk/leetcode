package leetcode

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
