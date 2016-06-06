package leetcode

import (
	"sort"
	"testing"
)

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	return intersectionNums(nums(nums1), nums(nums2))
}

func intersectionNums(n1, n2 nums) (n []int) {
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
        v := n1[i]
        n = append(n, v)
        for i++; i < len(n1); i++ {
            if n1[i] != v {
                break
            }
        }
        for j++; j < len(n2); j++ {
            if n2[j] != v {
                break
            }
        }
	}
	return
}

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

func TestIntersection(t *testing.T) {
	n1 := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	if len(n1) != 1 {
		t.Fatal("len(n1) != 1")
	}
	if n1[0] != 2 {
		t.Fatal(n1)
	}

	n2 := intersection([]int{4, 7, 9, 7, 6, 7}, []int{5, 0, 0, 6, 1, 6, 2, 2, 4})
	if len(n2) != 2 {
		t.Fatal("len(n2) != 2")
	}
	if n2[0] != 4 || n2[1] != 6 {
		t.Fatal(n2)
	}
}
