package _240_search_a_2d_matrix_ii

import (
	"fmt"
	"testing"
)

/**
240. 搜索二维矩阵 II
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false

提示：
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matix[i][j] <= 109
每行的所有元素从左到右升序排列
每列的所有元素从上到下升序排列
-109 <= target <= 109


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii

*/

/**
初始化指针，以左下角为起始点，值大于target指针向上移动，小于target指针向右移动
*/
func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{1,4,7,11,15},
		[]int{2,5,8,12,19},
		[]int{3,7,9,16,22},
		[]int{10,13,14,17,24},
		[]int{18,21,23,26,30},
	}

	fmt.Println(SearchMatrix(matrix, 5), SearchMatrix(matrix, 6))
}
