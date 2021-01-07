package _029_shun_shi_zhen_da_yin_ju_zhen_lcof

/**
剑指 Offer 29. 顺时针打印矩阵
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100

链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
*/

/**
时间复杂度：O(mn)，其中 m 和 n 分别是输入矩阵的行数和列数。矩阵中的每个元素都要被访问一次。
空间复杂度：O(1)。除了输出数组以外，空间复杂度是常数。
 */
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
