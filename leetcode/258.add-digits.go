// https://leetcode.com/problems/add-digits/

package leetcode

func addDigits(num int) int {
	var num2 int
	for num >= 10 {
		num2 += num / 10
		num = num % 10
	}
	if num2 += num; num2 >= 10 {
		return addDigits(num2)
	}
	return num2
}
