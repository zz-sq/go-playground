/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] Surrounded Regions
 *
 * https://leetcode.cn/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (46.57%)
 * Likes:    1131
 * Dislikes: 0
 * Total Accepted:    290.8K
 * Total Submissions: 624.3K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * You are given an m x n matrix board containing letters 'X' and 'O', capture
 * regions that are surrounded:
 *
 *
 * Connect: A cell is connected to adjacent cells horizontally or
 * vertically.
 * Region: To form a region connect every 'O' cell.
 * Surround: The region is surrounded with 'X' cells if you can connect the
 * region with 'X' cells and none of the region cells are on the edge of the
 * board.
 *
 *
 * A surrounded region is captured by replacing all 'O's with 'X's in the input
 * matrix board.
 *
 *
 * Example 1:
 *
 *
 * Input: board =
 * [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
 *
 * Output:
 * [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
 *
 * Explanation:
 *
 * In the above diagram, the bottom region is not captured because it is on the
 * edge of the board and cannot be surrounded.
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["X"]]
 *
 * Output: [["X"]]
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n == board[i].length
 * 1 <= m, n <= 200
 * board[i][j] is 'X' or 'O'.
 *
 *
 */

// @lc code=start
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0 || i == m-1 || j == n-1) && board[i][j] == 'O' {
				dfs(board, m, n, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, m, n, i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	board[i][j] = '#'
	dfs(board, m, n, i+1, j)
	dfs(board, m, n, i, j+1)
	dfs(board, m, n, i-1, j)
	dfs(board, m, n, i, j-1)
}

// @lc code=end

