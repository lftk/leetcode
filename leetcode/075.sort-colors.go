// https://leetcode.com/problems/sort-colors/

package leetcode

func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	for i := l; i <= r; i++ {
		for i < r && nums[i] > 1 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
		for i > l && nums[i] < 1 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
}
