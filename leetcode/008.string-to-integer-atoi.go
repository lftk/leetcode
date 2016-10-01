// https://leetcode.com/problems/string-to-integer-atoi/

package leetcode

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var (
		num int
		neg bool
	)
	for i, s := range str {
		if s != ' ' {
			str = str[i:]
			break
		}
	}
	if str[0] == '+' {
		str = str[1:]
	} else if str[0] == '-' {
		str = str[1:]
		neg = true
	}
	for _, s := range str {
		if s < '0' || s > '9' {
			break
		}
		num = num*10 + int(s-'0')
		if num > 2147483648 {
			num = 2147483648
			break
		}
	}
	if neg {
		num *= -1
	} else if num > 2147483647 {
		num = 2147483647
	}
	return num
}
