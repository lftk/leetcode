package leetcode

import "testing"

func Test_summaryRanges(t *testing.T) {
	s := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	if len(s) != 3 {
		t.Fatal(s)
	}
	if s[0] != "0->2" || s[1] != "4->5" || s[2] != "7" {
		t.Fatal(s)
	}
}
