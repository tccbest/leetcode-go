package _074_search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])
	left := 0
	right := m * n - 1

	for left <= right {
		mid := left + (right - left) >> 1
		if target == matrix[(mid - 1) / n][mid % n - 1] {
			return true
		} else if target < matrix[(mid - 1) / n][mid % n - 1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}