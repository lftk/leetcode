package leetcode

import "testing"

func Test_containsDuplicate(t *testing.T) {
	b := containsDuplicate([]int{1, 2, 3, 1, 4})
	if !b {
		t.Fatal("no found duplicate elements")
	}
}
