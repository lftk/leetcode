package leetcode

import "sort"

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
