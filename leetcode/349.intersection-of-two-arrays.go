package leetcode

import "sort"

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
