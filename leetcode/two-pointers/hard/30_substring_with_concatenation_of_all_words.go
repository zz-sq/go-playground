/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 *
 * https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (38.84%)
 * Likes:    1127
 * Dislikes: 0
 * Total Accepted:    212.3K
 * Total Submissions: 546.8K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * You are given a string s and an array of strings words. All the strings of
 * words are of the same length.
 *
 * A concatenated string is a string that exactly contains all the strings of
 * any permutation of words concatenated.
 *
 *
 * For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef",
 * "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is
 * not a concatenated string because it is not the concatenation of any
 * permutation of words.
 *
 *
 * Return an array of the starting indices of all the concatenated substrings
 * in s. You can return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "barfoothefoobarman", words = ["foo","bar"]
 *
 * Output: [0,9]
 *
 * Explanation:
 *
 * The substring starting at 0 is "barfoo". It is the concatenation of
 * ["bar","foo"] which is a permutation of words.
 * The substring starting at 9 is "foobar". It is the concatenation of
 * ["foo","bar"] which is a permutation of words.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "wordgoodgoodgoodbestword", words =
 * ["word","good","best","word"]
 *
 * Output: []
 *
 * Explanation:
 *
 * There is no concatenated substring.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
 *
 * Output: [6,9,12]
 *
 * Explanation:
 *
 * The substring starting at 6 is "foobarthe". It is the concatenation of
 * ["foo","bar","the"].
 * The substring starting at 9 is "barthefoo". It is the concatenation of
 * ["bar","the","foo"].
 * The substring starting at 12 is "thefoobar". It is the concatenation of
 * ["the","foo","bar"].
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * 1 <= words.length <= 5000
 * 1 <= words[i].length <= 30
 * s and words[i] consist of lowercase English letters.
 *
 *
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	mp := map[string]int{}
	for _, word := range words {
		mp[word]++
	}
	wordLen := len(words[0])
	res := []int{}
	for i := 0; i <= len(s)-len(words)*wordLen; i++ {
		visited := map[string]int{}
		for j := 0; j < len(words); j++ {
			k := i + j*wordLen
			word := s[k : k+wordLen]
			if _, ok := mp[word]; !ok {
				break
			}
			visited[word]++
			if visited[word] > mp[word] {
				break
			}
			if j == len(words)-1 {
				res = append(res, i)
			}
		}
	}
	return res
}

// @lc code=end

