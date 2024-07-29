/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (38.50%)
 * Likes:    7283
 * Dislikes: 0
 * Total Accepted:    1.8M
 * Total Submissions: 4.6M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, return the longest palindromic substring in s.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "babad"
 * Output: "bab"
 * Explanation: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "cbbd"
 * Output: "bb"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 1000
 * s consist of only digits and English letters.
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([]bool, n)
	var res string
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			dp[j] = s[i] == s[j] && (j-i < 3 || dp[j-1])
			if dp[j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

// @lc code=end

