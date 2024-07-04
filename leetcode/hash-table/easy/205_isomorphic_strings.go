/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] Isomorphic Strings
 *
 * https://leetcode.cn/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (49.35%)
 * Likes:    723
 * Dislikes: 0
 * Total Accepted:    280.6K
 * Total Submissions: 568.5K
 * Testcase Example:  '"egg"\n"add"'
 *
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings s and t are isomorphic if the characters in s can be replaced to
 * get t.
 *
 * All occurrences of a character must be replaced with another character while
 * preserving the order of characters. No two characters may map to the same
 * character, but a character may map to itself.
 *
 *
 * Example 1:
 * Input: s = "egg", t = "add"
 * Output: true
 * Example 2:
 * Input: s = "foo", t = "bar"
 * Output: false
 * Example 3:
 * Input: s = "paper", t = "title"
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 5 * 10^4
 * t.length == s.length
 * s and t consist of any valid ascii character.
 *
 *
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]byte, len(s))
	tMap := make(map[byte]byte, len(t))
	for i := 0; i < len(s); i++ {
		c, ok := sMap[s[i]]
		if ok {
			if c != t[i] {
				return false
			}
		} else {
			sMap[s[i]] = t[i]
		}
		c, ok = tMap[t[i]]
		if ok {
			if c != s[i] {
				return false
			}
		} else {
			tMap[t[i]] = s[i]
		}
	}
	return true
}

// @lc code=end

