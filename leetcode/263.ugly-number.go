// https://leetcode.com/problems/ugly-number/

package leetcode

func isUgly(num int) bool {
	factors := []int{2, 3, 5}
	for num > 1 {
		var factor int
		for _, v := range factors {
			if num >= v && num%v == 0 {
				factor = v
				break
			}
		}
		if factor == 0 {
			return false
		}
		num /= factor
	}
	return num == 1
}
