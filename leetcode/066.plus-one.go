// https://leetcode.com/problems/plus-one/

package leetcode

func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i] + c
		digits[i] = n % 10
		c = n / 10
	}
	if c == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
