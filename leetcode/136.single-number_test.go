// https://leetcode.com/problems/single-number/

package leetcode

func singleNumber(nums []int) int {
	var x int
	for _, v := range nums {
		x = x ^ v
	}
	return x
}
