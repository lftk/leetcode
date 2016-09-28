package leetcode

import (
    "testing"
)

func Test_isIsomorphic(t *testing.T) {
    b1 := isIsomorphic("aabbcc", "bbccaa")
    if !b1 {
        t.Error("aabbcc", "bbccaa")
    }

    b2 := isIsomorphic("aa", "ab")
    if !b2 {
        t.Error("aa", "ab")
    }
}
