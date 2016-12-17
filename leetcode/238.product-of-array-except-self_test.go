package leetcode

import (
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	vals := productExceptSelf([]int{1, 2, 3, 4})
	for i, val := range []int{24, 12, 8, 6} {
		if val != vals[i] {
			t.Error(i, val, vals[i])
		}
	}
}
