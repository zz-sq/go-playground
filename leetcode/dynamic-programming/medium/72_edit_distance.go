/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] Edit Distance
 *
 * https://leetcode.cn/problems/edit-distance/description/
 *
 * algorithms
 * Medium (62.96%)
 * Likes:    3428
 * Dislikes: 0
 * Total Accepted:    512.7K
 * Total Submissions: 814.4K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * Given two strings word1 and word2, return the minimum number of operations
 * required to convert word1 to word2.
 *
 * You have the following three operations permitted on a word:
 *
 *
 * Insert a character
 * Delete a character
 * Replace a character
 *
 *
 *
 * Example 1:
 *
 *
 * Input: word1 = "horse", word2 = "ros"
 * Output: 3
 * Explanation:
 * horse -> rorse (replace 'h' with 'r')
 * rorse -> rose (remove 'r')
 * rose -> ros (remove 'e')
 *
 *
 * Example 2:
 *
 *
 * Input: word1 = "intention", word2 = "execution"
 * Output: 5
 * Explanation:
 * intention -> inention (remove 't')
 * inention -> enention (replace 'i' with 'e')
 * enention -> exention (replace 'n' with 'x')
 * exention -> exection (replace 'n' with 'c')
 * exection -> execution (insert 'u')
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= word1.length, word2.length <= 500
 * word1 and word2 consist of lowercase English letters.
 *
 *
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

// @lc code=end

