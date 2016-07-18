// https://leetcode.com/problems/candy/

package leetcode

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	candys := make([]int, n)
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] && candys[i] <= candys[i-1] {
			candys[i] = candys[i-1] + 1
		}
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candys[i-1] <= candys[i] {
			candys[i-1] = candys[i] + 1
		}
	}
	for _, c := range candys {
		n += c
	}
	return n
}
