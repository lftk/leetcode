package leetcode

import "testing"

func Test_moveZeroes(t *testing.T) {
    n := moveZeroes([]int32{0, 1, 0, 3, 12})
    if len(n) != 5 || n[0] != 1 || n[1] != 3 || n[2] != 12 || n[3] != 0 || n[4] != 0 {
        t.Fatal(n)
    }
}
