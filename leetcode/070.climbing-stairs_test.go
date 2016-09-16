package leetcode

import (
	"testing"
)

func Test_climbStairs(t *testing.T) {
	n := climbStairs(30)
	if n != 1346269 {
		t.Error(n, 30)
	}
}
