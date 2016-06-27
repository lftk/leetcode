// https://leetcode.com/problems/majority-element/

package leetcode

func majorityElement(nums []int) int {
    m, c := nums[0], 1
    for _, n := range nums[1:] {
        if m == n {
            c++
        } else {
            if c--; c == 0 {
                m = n
                c = 1
            }
        }
    }
    return m
}
