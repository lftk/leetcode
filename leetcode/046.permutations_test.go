package leetcode

import "testing"

func Test_permute(t *testing.T) {
	ret := permute([]int{1, 2, 3})
	if len(ret) != 6 {
		t.Fatal(ret)
	}
	if r := ret[0]; r[0] != 1 || r[1] != 2 || r[2] != 3 {
		t.Fatal(ret)
	}
	if r := ret[5]; r[0] != 3 || r[1] != 2 || r[2] != 1 {
		t.Fatal(ret)
	}
}
