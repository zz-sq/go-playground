/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (51.12%)
 * Likes:    1708
 * Dislikes: 0
 * Total Accepted:    543.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given an m x n matrix, return all elements of the matrix in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	var res []int
	dfs(matrix, &res, 0, 0)
	return res
}

func dfs(matrix [][]int, res *[]int, i int, j int) {
	if i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[0]) && matrix[i][j] != 1000 {
		*res = append(*res, matrix[i][j])
		matrix[i][j] = 1000
		if i == 0 || matrix[i-1][j] == 1000 {
			dfs(matrix, res, i, j+1)
		}
		dfs(matrix, res, i+1, j)
		dfs(matrix, res, i, j-1)
		dfs(matrix, res, i-1, j)
	}
}

// @lc code=end

