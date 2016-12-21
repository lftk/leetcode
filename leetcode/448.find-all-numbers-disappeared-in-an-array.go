// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

package leetcode

func findDisappearedNumbers(nums []int) []int {
	var rst []int
	for _, num := range nums {
		val := num
		if num < 0 {
			val = -num
		}
		if val--; nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}
	for i, num := range nums {
		if num > 0 {
			rst = append(rst, i+1)
		}
	}
	return rst
}
