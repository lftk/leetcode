// https://leetcode.com/problems/house-robber/

package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	vals := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			vals[i] = num
		} else if i == 1 {
			vals[i] = max(num, vals[i-1])
		} else {
			vals[i] = max(num+vals[i-2], vals[i-1])
		}
	}
	return vals[len(nums)-1]
}
