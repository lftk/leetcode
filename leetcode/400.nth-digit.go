// https://leetcode.com/problems/nth-digit/

package leetcode

func findNthDigit(n int) int {
	d, base := 1, 1
	for n > 9*base*d {
		n -= 9 * base * d
		base *= 10
		d++
	}
	num := (n-1)/d + base
	for i := 1; i < d-(n-1)%d; i++ {
		num /= 10
	}
	return num % 10
}
