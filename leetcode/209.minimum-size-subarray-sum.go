// https://leetcode.com/problems/minimum-size-subarray-sum/

package leetcode

func minSubArrayLen(s int, nums []int) int {
	var i, j, t, n int
	for j < len(nums) {
		for t < s && j < len(nums) {
			t += nums[j]
			j++
		}
		for t >= s && i < len(nums) {
			if n == 0 || n > j-i {
				n = j - i
			}
			t -= nums[i]
			i++
		}
	}
	return n
}
