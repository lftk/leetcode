// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max, n0 := 0, prices[0]
	for _, n := range prices[1:] {
		if n > n0 {
			if max < n-n0 {
				max = n - n0
			}
		} else {
			n0 = n
		}
	}
	return max
}
