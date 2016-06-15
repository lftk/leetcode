package leetcode

import "testing"

func Test_searchMatrix(t *testing.T) {
	m := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	if b := searchMatrix(m, 3); !b {
		t.Fatal("no find 3")
	}
	if b := searchMatrix(m, 8); b {
		t.Fatal("find 8")
	}
}
