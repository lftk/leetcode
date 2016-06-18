// https://leetcode.com/problems/single-number-ii/

package leetcode

func singleNumber2(nums []int) int {
	var one, two int
	for _, n := range nums {
		two |= one & n
		one ^= n
		three := one & two
		one &= ^three
		two &= ^three
	}
	return one
}
