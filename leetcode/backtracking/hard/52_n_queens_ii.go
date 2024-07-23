/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N-Queens II
 *
 * https://leetcode.cn/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (82.33%)
 * Likes:    519
 * Dislikes: 0
 * Total Accepted:    154.9K
 * Total Submissions: 188.1K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n x n
 * chessboard such that no two queens attack each other.
 *
 * Given an integer n, return the number of distinct solutions to the n-queens
 * puzzle.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4
 * Output: 2
 * Explanation: There are two distinct solutions to the 4-queens puzzle as
 * shown.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 9
 *
 *
 */

// @lc code=start
func totalNQueens(n int) int {
	res := 0
	rows := make([]bool, n)
	c1 := make([]bool, 2*n)
	c2 := make([]bool, 2*n)
	var dfs func(k int)
	dfs = func(k int) {
		if k == n {
			res++
			return
		}
		for i := 0; i < n; i++ {
			idx1, idx2 := k-i+n-1, k+i
			if rows[i] || c1[idx1] || c2[idx2] {
				continue
			}
			rows[i], c1[idx1], c2[idx2] = true, true, true
			dfs(k + 1)
			rows[i], c1[idx1], c2[idx2] = false, false, false
		}
	}
	dfs(0)
	return res
}

// @lc code=end

