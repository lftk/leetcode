// https://leetcode.com/problems/hamming-distance/

package leetcode

func hammingDistance(x int, y int) int {
	var n int
	z := x ^ y
	for z != 0 {
		z = z & (z - 1)
		n++
	}
	return n
}
