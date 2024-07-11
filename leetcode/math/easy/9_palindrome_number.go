/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.cn/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (56.12%)
 * Likes:    2864
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 2.9M
 * Testcase Example:  '121'
 *
 * Given an integer x, return true if x is a palindrome, and false
 * otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: x = 121
 * Output: true
 * Explanation: 121 reads as 121 from left to right and from right to left.
 *
 *
 * Example 2:
 *
 *
 * Input: x = -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: x = 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= x <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you solve it without converting the integer to a string?
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	res, n := 0, x
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res == n
}

// @lc code=end

