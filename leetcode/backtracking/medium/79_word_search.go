/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.cn/problems/word-search/description/
 *
 * algorithms
 * Medium (47.21%)
 * Likes:    1844
 * Dislikes: 0
 * Total Accepted:    546.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given an m x n grid of characters board and a string word, return true if
 * word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cells,
 * where adjacent cells are horizontally or vertically neighboring. The same
 * letter cell may not be used more than once.
 *
 *
 * Example 1:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCCED"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "SEE"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCB"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n = board[i].length
 * 1 <= m, n <= 6
 * 1 <= word.length <= 15
 * board and word consists of only lowercase and uppercase English letters.
 *
 *
 *
 * Follow up: Could you use search pruning to make your solution faster with a
 * larger board?
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	res := false
	var dfs func(k int, i int, j int)
	dfs = func(k int, i int, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] == '#' || res {
			return
		}
		if board[i][j] == word[k] {
			if k == len(word)-1 {
				res = true
				return
			}
			board[i][j] = '#'
			dfs(k+1, i+1, j)
			dfs(k+1, i-1, j)
			dfs(k+1, i, j+1)
			dfs(k+1, i, j-1)
			board[i][j] = word[k]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(0, i, j)
		}
	}
	return res
}

// @lc code=end

