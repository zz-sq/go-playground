/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] Is Subsequence
 *
 * https://leetcode.cn/problems/is-subsequence/description/
 *
 * algorithms
 * Easy (52.51%)
 * Likes:    1053
 * Dislikes: 0
 * Total Accepted:    446.6K
 * Total Submissions: 850.4K
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * Given two strings s and t, return true if s is a subsequence of t, or false
 * otherwise.
 *
 * A subsequence of a string is a new string that is formed from the original
 * string by deleting some (can be none) of the characters without disturbing
 * the relative positions of the remaining characters. (i.e., "ace" is a
 * subsequence of "abcde" while "aec" is not).
 *
 *
 * Example 1:
 * Input: s = "abc", t = "ahbgdc"
 * Output: true
 * Example 2:
 * Input: s = "axc", t = "ahbgdc"
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 100
 * 0 <= t.length <= 10^4
 * s and t consist only of lowercase English letters.
 *
 *
 *
 * Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k
 * >= 10^9, and you want to check one by one to see if t has its subsequence.
 * In this scenario, how would you change your code?
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

// @lc code=end

