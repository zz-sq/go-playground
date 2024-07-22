/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.cn/problems/word-ladder/description/
 *
 * algorithms
 * Hard (48.95%)
 * Likes:    1385
 * Dislikes: 0
 * Total Accepted:    222.2K
 * Total Submissions: 453.9K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * A transformation sequence from word beginWord to word endWord using a
 * dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... ->
 * sk such that:
 *
 *
 * Every adjacent pair of words differs by a single letter.
 * Every si for 1 <= i <= k is in wordList. Note that beginWord does not need
 * to be in wordList.
 * sk == endWord
 *
 *
 * Given two words, beginWord and endWord, and a dictionary wordList, return
 * the number of words in the shortest transformation sequence from beginWord
 * to endWord, or 0 if no such sequence exists.
 *
 *
 * Example 1:
 *
 *
 * Input: beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log","cog"]
 * Output: 5
 * Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot"
 * -> "dog" -> cog", which is 5 words long.
 *
 *
 * Example 2:
 *
 *
 * Input: beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log"]
 * Output: 0
 * Explanation: The endWord "cog" is not in wordList, therefore there is no
 * valid transformation sequence.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= beginWord.length <= 10
 * endWord.length == beginWord.length
 * 1 <= wordList.length <= 5000
 * wordList[i].length == beginWord.length
 * beginWord, endWord, and wordList[i] consist of lowercase English
 * letters.
 * beginWord != endWord
 * All the words in wordList are unique.
 *
 *
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	type Pair struct {
		Word string
		Step int
	}
	queue := []Pair{{beginWord, 0}}
	visted := map[string]bool{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if distance(cur.Word, endWord) == 0 {
			return cur.Step + 1
		}
		for _, candidate := range wordList {
			if distance(cur.Word, candidate) == 1 && !visted[candidate] {
				visted[candidate] = true
				queue = append(queue, Pair{candidate, cur.Step + 1})
			}
		}
	}
	return 0
}

func distance(m, n string) int {
	d := 0
	for i := 0; i < len(m); i++ {
		if m[i] != n[i] {
			d++
		}
	}
	return d
}

// @lc code=end

