package leetcode

import "testing"

func Test_diffWaysToCompute(t *testing.T) {
	ret := diffWaysToCompute("2*3-4*5")
	if len(ret) != 5 {
		t.Fatal("len(ret) != 5")
	}

	if !compareSlice(ret, []int{-34, -10, -14, -10, 10}) {
		t.Fatal(ret)
	}
}
