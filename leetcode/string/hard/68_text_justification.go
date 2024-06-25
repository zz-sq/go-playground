/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] Text Justification
 *
 * https://leetcode.cn/problems/text-justification/description/
 *
 * algorithms
 * Hard (54.09%)
 * Likes:    427
 * Dislikes: 0
 * Total Accepted:    80.8K
 * Total Submissions: 149.4K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * Given an array of strings words and a width maxWidth, format the text such
 * that each line has exactly maxWidth characters and is fully (left and right)
 * justified.
 *
 * You should pack your words in a greedy approach; that is, pack as many words
 * as you can in each line. Pad extra spaces ' ' when necessary so that each
 * line has exactly maxWidth characters.
 *
 * Extra spaces between words should be distributed as evenly as possible. If
 * the number of spaces on a line does not divide evenly between words, the
 * empty slots on the left will be assigned more spaces than the slots on the
 * right.
 *
 * For the last line of text, it should be left-justified, and no extra space
 * is inserted between words.
 *
 * Note:
 *
 *
 * A word is defined as a character sequence consisting of non-space characters
 * only.
 * Each word's length is guaranteed to be greater than 0 and not exceed
 * maxWidth.
 * The input array words contains at least one word.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: words = ["This", "is", "an", "example", "of", "text",
 * "justification."], maxWidth = 16
 * Output:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 *
 * Example 2:
 *
 *
 * Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth
 * = 16
 * Output:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * Explanation: Note that the last line is "shall be    " instead of "shall
 * be", because the last line must be left-justified instead of
 * fully-justified.
 * Note that the second line is also left-justified because it contains only
 * one word.
 *
 * Example 3:
 *
 *
 * Input: words =
 * ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"],
 * maxWidth = 20
 * Output:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 *
 *
 * Constraints:
 *
 *
 * 1 <= words.length <= 300
 * 1 <= words[i].length <= 20
 * words[i] consists of only English letters and symbols.
 * 1 <= maxWidth <= 100
 * words[i].length <= maxWidth
 *
 *
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	ret := make([][]string, len(words))
	wc := make([]int, 0)
	i, curLen := 0, 0
	for _, word := range words {
		if len(ret[i])+curLen+len(word) <= maxWidth {
			ret[i] = append(ret[i], word)
			curLen += len(word)
		} else {
			wc = append(wc, curLen)
			curLen = len(word)
			i++
			ret[i] = append(ret[i], word)
		}
	}
	wc = append(wc, curLen)
	res := make([]string, 0)
	for m := 0; m <= i; m++ {
		line := ret[m]
		leftSpace := maxWidth - wc[m]
		var perWord, extraSpace int
		if len(line) == 1 {
			perWord = leftSpace
			extraSpace = 0
		} else if m == i {
			perWord = 1
			extraSpace = 0
		} else {
			perWord = leftSpace / (len(line) - 1)
			extraSpace = leftSpace % (len(line) - 1)
		}
		var builder strings.Builder
		for j, l := range line {
			builder.WriteString(l)
			if j != 0 && j == len(line)-1 {
				continue
			}
			for k := 0; k < perWord; k++ {
				builder.WriteRune(' ')
			}
			if j < extraSpace {
				builder.WriteRune(' ')
			}
		}
		if m == i {
			len := builder.Len()
			for k := 0; k < maxWidth-len; k++ {
				builder.WriteRune(' ')
			}
		}
		res = append(res, builder.String())
	}
	return res
}

// @lc code=end

