package leetcode

import "testing"

func Test_searchRange(t *testing.T) {
	r := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	if r[0] != 3 || r[1] != 4 {
		t.Fatal(r)
	}
}
