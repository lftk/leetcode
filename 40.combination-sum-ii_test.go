package leetcode

import (
	"sort"
	"testing"
)

func combinationSum2(candidates []int, target int) [][]int {
	c := Candidates(candidates)
	sort.Sort(c)
	return orderedCombinationSum2(c, target)
}

type Candidates []int

func (c Candidates) Len() int {
	return len(c)
}

func (c Candidates) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Candidates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func orderedCombinationSum2(c Candidates, target int) [][]int {
	var ret [][]int
	for i := 0; i < c.Len(); i++ {
		n := c[i]
		if n > target {
			break
		} else if n == target {
			ret = append(ret, []int{n})
		} else if c.Len() != 1 {
			if s := orderedCombinationSum2(c[i+1:], target-n); s != nil {
				for j := 0; j < len(s); j++ {
					s[j] = append([]int{n}, s[j]...)
					ret = append(ret, s[j])
				}
			}
		}

		for ; i+1 < c.Len(); i++ {
			if n != c[i+1] {
				break
			}
		}
	}
	return ret
}

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

func TestCombinationSum2(t *testing.T) {
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
