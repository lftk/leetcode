// https://leetcode.com/problems/isomorphic-strings/

package leetcode

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    ss, tt := make(map[byte]int), make(map[byte]int)
    for i, j := 0, 0; i < len(s); i, j= i+1, j+1 {
        b1, b2 := s[i], t[j]
        if ss[b1] != tt[b2] {
            return false
        }
        ss[b1], tt[b2] = i + 1, j + 1
    }
    return true
}
