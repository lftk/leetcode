package leetcode

import "testing"

func compareSlice(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func Test_combinationSum2(t *testing.T) {
	s := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)

	if len(s) != 4 {
		t.Fatalf("len(s1) != 4")
	}

	if !compareSlice(s[0], []int{1, 1, 6}) {
		t.Fatal(s[0])
	}

	if !compareSlice(s[1], []int{1, 2, 5}) {
		t.Fatal(s[1])
	}

	if !compareSlice(s[2], []int{1, 7}) {
		t.Fatal(s[2])
	}

	if !compareSlice(s[3], []int{2, 6}) {
		t.Fatal(s[3])
	}
}
