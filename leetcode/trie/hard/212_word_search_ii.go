/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] Word Search II
 *
 * https://leetcode.cn/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (43.15%)
 * Likes:    878
 * Dislikes: 0
 * Total Accepted:    110Kp
 * Total Submissions: 254.8K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
  '["oath","pea","eat","rain"]'
 *
 * Given an m x n board of characters and a list of strings words, return all
 * words on the board.
 *
 * Each word must be constructed from letters of sequentially adjacent cells,
 * where adjacent cells are horizontally or vertically neighboring. The same
 * letter cell may not be used more than once in a word.
 *
 *
 * Example 1:
 *
 *
 * Input: board =
 * [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
 * words = ["oath","pea","eat","rain"]
 * Output: ["eat","oath"]
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["a","b"],["c","d"]], words = ["abcb"]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n == board[i].length
 * 1 <= m, n <= 12
 * board[i][j] is a lowercase English letter.
 * 1 <= words.length <= 3 * 10^4
 * 1 <= words[i].length <= 10
 * words[i] consists of lowercase English letters.
 * All the strings of words are unique.
 *
 *
*/

// @lc code=start
type Trie struct {
	children [26]*Trie
	isWord   bool
	word     string
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, ch := range word {
		idx := ch - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &Trie{}
		}
		cur = cur.children[idx]
	}
	cur.isWord = true
	cur.word = word
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return res
	}
	m, n := len(board), len(board[0])
	root := &Trie{}
	for _, word := range words {
		root.Insert(word)
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(node *Trie, i int, j int, visited [][]bool)
	dfs = func(node *Trie, i int, j int, visited [][]bool) {
		if i < 0 || j < 0 || i >= m || j >= n || visited[i][j] {
			return
		}
		idx := board[i][j] - 'a'
		next := node.children[idx]
		if next == nil {
			return
		}
		if next.isWord {
			res = append(res, next.word)
			next.isWord = false
		}
		visited[i][j] = true
		dfs(next, i+1, j, visited)
		dfs(next, i-1, j, visited)
		dfs(next, i, j+1, visited)
		dfs(next, i, j-1, visited)
		visited[i][j] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(root, i, j, visited)
		}
	}
	return res
}

// @lc code=end

