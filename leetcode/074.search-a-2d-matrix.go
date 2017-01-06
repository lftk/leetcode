package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	start, end, mid := 0, m*n-1, 0
	for start <= end {
		mid = start + (end-start)/2
		v := matrix[mid/n][mid%n]

		if target < v {
			end = mid - 1
		} else if target > v {
			start = mid + 1
		} else {
			return true
		}
	}

	return false
}
