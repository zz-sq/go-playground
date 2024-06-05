/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] Ransom Note
 *
 * https://leetcode.cn/problems/ransom-note/description/
 *
 * algorithms
 * Easy (64.43%)
 * Likes:    884
 * Dislikes: 0
 * Total Accepted:    496.1K
 * Total Submissions: 769.7K
 * Testcase Example:  '"a"\n"b"'
 *
 * Given two strings ransomNote and magazine, return true if ransomNote can be
 * constructed by using the letters from magazine and false otherwise.
 *
 * Each letter in magazine can only be used once in ransomNote.
 *
 *
 * Example 1:
 * Input: ransomNote = "a", magazine = "b"
 * Output: false
 * Example 2:
 * Input: ransomNote = "aa", magazine = "ab"
 * Output: false
 * Example 3:
 * Input: ransomNote = "aa", magazine = "aab"
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= ransomNote.length, magazine.length <= 10^5
 * ransomNote and magazine consist of lowercase English letters.
 *
 *
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	var count = make(map[rune]int)
	for _, v := range magazine {
		count[v] += 1
	}
	for _, v := range ransomNote {
		count[v] -= 1
		if count[v] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

