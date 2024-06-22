/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.cn/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (45.87%)
 * Likes:    707
 * Dislikes: 0
 * Total Accepted:    584.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consisting of words and spaces, return the length of the
 * last word in the string.
 *
 * A word is a maximal substring consisting of non-space characters only.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "Hello World"
 * Output: 5
 * Explanation: The last word is "World" with length 5.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "   fly me   to   the moon  "
 * Output: 4
 * Explanation: The last word is "moon" with length 4.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "luffy is still joyboy"
 * Output: 6
 * Explanation: The last word is "joyboy" with length 6.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of only English letters and spaces ' '.
 * There will be at least one word in s.
 *
 *
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	res, start := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		v := s[i]
		if v == ' ' {
			if start == 0 {
				continue
			} else {
				return res
			}
		} else if start == 0 {
			start = 1
			res++
		} else {
			res++
		}
	}
	return res
}

// @lc code=end

