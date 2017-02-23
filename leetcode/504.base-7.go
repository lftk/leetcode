// https://leetcode.com/problems/base-7

package leetcode

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var (
		s        string
		negative bool
	)
	if num < 0 {
		num = -num
		negative = true
	}
	for num > 0 {
		n := num % 7
		s = string('0'+n) + s
		num /= 7
	}
	if negative {
		s = "-" + s
	}
	return s
}
