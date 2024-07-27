/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] Word Break
 *
 * https://leetcode.cn/problems/word-break/description/
 *
 * algorithms
 * Medium (55.96%)
 * Likes:    2543
 * Dislikes: 0
 * Total Accepted:    617.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * Given a string s and a dictionary of strings wordDict, return true if s can
 * be segmented into a space-separated sequence of one or more dictionary
 * words.
 *
 * Note that the same word in the dictionary may be reused multiple times in
 * the segmentation.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetcode", wordDict = ["leet","code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet
 * code".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "applepenapple", wordDict = ["apple","pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple
 * pen apple".
 * Note that you are allowed to reuse a dictionary word.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 300
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 20
 * s and wordDict[i] consist of only lowercase English letters.
 * All the strings of wordDict are unique.
 *
 *
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := map[int]bool{}
	dp[-1] = true
	for i := 0; i < len(s); i++ {
		for j, _ := range dp {
			cur := s[j+1 : i+1]
			fmt.Println(cur)
			if wordMap[cur] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)-1]
}

// @lc code=end

