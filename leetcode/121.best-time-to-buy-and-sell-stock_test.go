package leetcode

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	n1 := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if n1 != 5 {
		t.Error(n1, 5)
	}
}
