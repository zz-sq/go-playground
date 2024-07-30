/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] Maximal Square
 *
 * https://leetcode.cn/problems/maximal-square/description/
 *
 * algorithms
 * Medium (50.53%)
 * Likes:    1678
 * Dislikes: 0
 * Total Accepted:    336.5K
 * Total Submissions: 665.9K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * Given an m x n binary matrix filled with 0's and 1's, find the largest
 * square containing only 1's and return its area.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [["0","1"],["1","0"]]
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: matrix = [["0"]]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 300
 * matrix[i][j] is '0' or '1'.
 *
 *
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	res, m, n := 0, len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			res = max(res, dp[i][0])
		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			res = max(res, dp[0][i])
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			res = max(res, dp[i][j])
		}
	}
	return res * res
}

// @lc code=end

