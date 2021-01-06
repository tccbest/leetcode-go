package _054_spiral_matrix

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	row := len(matrix)    //列长度
	col := len(matrix[0]) //行长度

	ret := make([]int, row*col)                    //返回值
	top, left, right, bottom := 0, 0, col-1, row-1 //上，左，右，下	左上角00
	index := 0                                     //当前位置索引

	for top <= bottom && left <= right {
		//从左上角遍历到右上角
		for colIndex := left; colIndex <= right; colIndex++ {
			ret[index] = matrix[top][colIndex]
			index++
		}

		//从右上角遍历到右下角
		for colIndex := top + 1; colIndex <= bottom; colIndex++ {
			ret[index] = matrix[colIndex][right]
			index++
		}

		if left < right && top < bottom {
			//从右下角遍历到左下角
			for colIndex := right - 1; colIndex > left; colIndex-- {
				ret[index] = matrix[bottom][colIndex]
				index++
			}

			//从左下角遍历到左上角
			for rowIndex := bottom; rowIndex > top; rowIndex-- {
				ret[index] = matrix[rowIndex][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return ret
}
