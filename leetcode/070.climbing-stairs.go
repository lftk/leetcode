// https://leetcode.com/problems/climbing-stairs/

package leetcode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		n2, n1 = n1+n2, n2
	}
	return n2
}
