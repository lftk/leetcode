// https://leetcode.com/problems/powx-n/

package leetcode

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	var ret float64
	if n < 0 {
		ret = 1 / myPow(x, -n)
	} else if n == 0 {
		ret = 1
	} else if n == 1 {
		ret = x
	} else if n%2 == 0 {
		ret = myPow(x, n/2)
		ret *= ret
	} else {
		ret = myPow(x, n/2)
		ret *= ret * x
	}
	return ret
}
