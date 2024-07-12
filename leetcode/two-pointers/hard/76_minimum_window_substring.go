/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (45.94%)
 * Likes:    2933
 * Dislikes: 0
 * Total Accepted:    595.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given two strings s and t of lengths m and n respectively, return the
 * minimum window substring of s such that every character in t (including
 * duplicates) is included in the window. If there is no such substring, return
 * the empty string "".
 *
 * The testcases will be generated such that the answer is unique.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C'
 * from string t.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "a", t = "a"
 * Output: "a"
 * Explanation: The entire string s is the minimum window.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "a", t = "aa"
 * Output: ""
 * Explanation: Both 'a's from t must be included in the window.
 * Since the largest window of s only has one 'a', return empty string.
 *
 *
 *
 * Constraints:
 *
 *
 * m == s.length
 * n == t.length
 * 1 <= m, n <= 10^5
 * s and t consist of uppercase and lowercase English letters.
 *
 *
 *
 * Follow up: Could you find an algorithm that runs in O(m + n) time?
 *
 */

// @lc code=start
func minWindow(s string, t string) string {
	needs := map[byte]int{}
	current := map[byte]int{}
	for _, c := range t {
		needs[byte(c)]++
	}
	valid, l, r, start, end := 0, 0, 0, -1, len(s)+1
	for r < len(s) {
		if needs[s[r]] > 0 {
			current[s[r]]++
			if current[s[r]] == needs[s[r]] {
				valid++
			}
		}
		for valid == len(needs) {
			if end-start > r-l {
				start = l
				end = r + 1
			}
			if needs[s[l]] > 0 {
				if needs[s[l]] == current[s[l]] {
					valid--
				}
				current[s[l]]--
			}
			l++
		}
		r++
	}
	if start == -1 {
		return ""
	} else {
		return s[start:end]
	}
}

// @lc code=end

