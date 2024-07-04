/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] Word Pattern
 *
 * https://leetcode.cn/problems/word-pattern/description/
 *
 * algorithms
 * Easy (44.70%)
 * Likes:    654
 * Dislikes: 0
 * Total Accepted:    188.5K
 * Total Submissions: 421.6K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * Given a pattern and a string s, find if s follows the same pattern.
 *
 * Here follow means a full match, such that there is a bijection between a
 * letter in pattern and a non-empty word in s.
 *
 *
 * Example 1:
 *
 *
 * Input: pattern = "abba", s = "dog cat cat dog"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: pattern = "abba", s = "dog cat cat fish"
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: pattern = "aaaa", s = "dog cat cat dog"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= pattern.length <= 300
 * pattern contains only lower-case English letters.
 * 1 <= s.length <= 3000
 * s contains only lowercase English letters and spaces ' '.
 * s does not contain any leading or trailing spaces.
 * All the words in s are separated by a single space.
 *
 *
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	word := strings.Split(s, " ")
	if len(pattern) != len(word) {
		return false
	}
	patternMaps := make(map[rune]string)
	wordMaps := make(map[string]rune)
	for i, p := range pattern {
		if c, ok := patternMaps[p]; ok {
			if c != word[i] {
				return false
			}
		} else {
			patternMaps[p] = word[i]
		}
		if c, ok := wordMaps[word[i]]; ok {
			if c != p {
				return false
			}
		} else {
			wordMaps[word[i]] = p
		}
	}
	return true
}

// @lc code=end

