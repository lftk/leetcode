// https://leetcode.com/problems/happy-number/

package leetcode

func isHappy(n int) bool {
	return isHappyInner(n, nil)
}

func isHappyInner(n int, t []int) bool {
	if n == 1 {
		return true
	}
	t = append(t, n)
	var nn int
	for n > 0 {
		i := n % 10
		nn += i * i
		n /= 10
	}
	for _, v := range t {
		if v == nn {
			return false
		}
	}
	return isHappyInner(nn, t)
}
