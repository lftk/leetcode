// https://leetcode.com/problems/bitwise-and-of-numbers-range/

package leetcode

func rangeBitwiseAnd(m int, n int) int {
	for n > m {
		n &= (n - 1)
	}
	return n
}
