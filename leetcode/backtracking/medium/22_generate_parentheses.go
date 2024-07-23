/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.86%)
 * Likes:    3626
 * Dislikes: 0
 * Total Accepted:    871.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '3'
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 * Example 1:
 * Input: n = 3
 * Output: ["((()))","(()())","(())()","()(())","()()()"]
 * Example 2:
 * Input: n = 1
 * Output: ["()"]
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(l int, r int, cur []rune)
	dfs = func(l int, r int, cur []rune) {
		if l == n && r == n {
			res = append(res, string(cur))
		}
		if l < n {
			cur = append(cur, '(')
			dfs(l+1, r, cur)
			cur = cur[:len(cur)-1]
		}
		if l > r {
			cur = append(cur, ')')
			dfs(l, r+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, 0, []rune{})
	return res
}

// @lc code=end

