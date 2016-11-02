// https://leetcode.com/problems/compare-version-numbers/

package leetcode

func compareVersion(version1 string, version2 string) int {
	atoi := func(str string) (n int) {
		for _, s := range str {
			n = n*10 + int(s-'0')
		}
		return n
	}

	parse := func(str string) (ver []int) {
		var i, j int
		for str += "."; j < len(str); i = j {
			for str[j] != '.' {
				j++
			}
			ver = append(ver, atoi(str[i:j]))
			j++
		}
		return ver
	}

	fill := func(ver []int, n int) []int {
		for i := len(ver); i < n; i++ {
			ver = append(ver, 0)
		}
		return ver
	}

	ver1 := parse(version1)
	ver2 := parse(version2)

	if len(ver1) < len(ver2) {
		ver1 = fill(ver1, len(ver2))
	} else if len(ver1) > len(ver2) {
		ver2 = fill(ver2, len(ver1))
	}

	for i, v1 := range ver1 {
		v2 := ver2[i]
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}
	return 0
}
