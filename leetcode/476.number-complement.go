// https://leetcode.com/problems/number-complement/

package leetcode

func findComplement(num int) int {
	var val int
	for n := uint(0); num != 0; num >>= 1 {
		val |= (1 - (num & 1)) << n
		n++
	}
	return val
}
