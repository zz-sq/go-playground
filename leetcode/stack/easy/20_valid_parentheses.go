/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.04%)
 * Likes:    4475
 * Dislikes: 0
 * Total Accepted:    1.9M
 * Total Submissions: 4.2M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Every close bracket has a corresponding open bracket of the same type.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s = "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: s = "(]"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 *
 *
 */

// @lc code=start
func isValid(s string) bool {
	pair := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}
	for _, v := range s {
		if _, ok := pair[v]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != pair[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}

// @lc code=end

