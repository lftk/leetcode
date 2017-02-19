// https://leetcode.com/problems/keyboard-row

package leetcode

func findWords(words []string) []string {
	alphabets := map[rune]int{
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}

	var dstWords []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		v := 0
		b := true
		for i, w := range word {
			if w >= 'A' && w <= 'Z' {
				w = w + 'a' - 'A'
			}
			if n := alphabets[w]; i == 0 {
				v = n
			} else if n != v {
				b = false
				break
			}
		}
		if b {
			dstWords = append(dstWords, word)
		}
	}
	return dstWords
}
