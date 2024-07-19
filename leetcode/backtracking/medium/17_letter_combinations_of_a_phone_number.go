/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (60.03%)
 * Likes:    2851
 * Dislikes: 0
 * Total Accepted:    910K
 * Total Submissions: 1.5M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent. Return the answer in
 * any order.
 *
 * A mapping of digits to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 * Example 1:
 *
 *
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * Example 2:
 *
 *
 * Input: digits = ""
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: digits = "2"
 * Output: ["a","b","c"]
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] is a digit in the range ['2', '9'].
 *
 *
 */

// @lc code=start
func letterCombinations(digits string) []string {
	maps := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	var backstrace func(i int, cur []rune)
	backstrace = func(i int, cur []rune) {
		if i == len(digits) {
			res = append(res, string(cur))
			return
		}
		letters := maps[digits[i]]
		for _, letter := range letters {
			cur = append(cur, letter)
			backstrace(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	backstrace(0, []rune{})
	return res
}

// @lc code=end

