package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	k, l := -1, len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		s := matrix[i][0]
		e := matrix[i][l-1]
		if s <= target && e >= target {
			k = i
			break
		}
	}
	if k == -1 {
		return false
	}

	l, r := 0, l
	for l <= r {
		m := (l + r) / 2
		v := matrix[k][m]
		if v == target {
			return true
		}
		if v > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}
