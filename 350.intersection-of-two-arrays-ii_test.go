package leetcode

import (
	"sort"
	"testing"
)

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	return intersectNums(nums(nums1), nums(nums2))
}

func intersectNums(n1, n2 nums) (n []int) {
	sort.Sort(n1)
	sort.Sort(n2)
	i, j := 0, 0
	for i < len(n1) && j < len(n2) {
        for n1[i] != n2[j] {
            for n1[i] < n2[j] {
                if i++; i == len(n1) {
                    return
                }
            }
            for n1[i] > n2[j] {
                if j++; j == len(n2) {
                    return
                }
            }
        }
        n = append(n, n1[i])
        i++
        j++
	}
	return
}

/*
type nums []int

func (n nums) Len() int {
	return len(n)
}

func (n nums) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
*/

func TestIntersect(t *testing.T) {
	n1 := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	if len(n1) != 2 {
		t.Fatal("len(n1) != 2")
	}
	if n1[0] != 2 && n1[1] != 2 {
		t.Fatal(n1)
	}

	n2 := intersect([]int{4, 7, 9, 7, 6, 7}, []int{5, 0, 0, 6, 1, 6, 2, 2, 4})
	if len(n2) != 2 {
		t.Fatal("len(n2) != 2")
	}
	if n2[0] != 4 || n2[1] != 6 {
		t.Fatal(n2)
	}
}
