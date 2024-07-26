/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 *
 * https://leetcode.cn/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (54.78%)
 * Likes:    521
 * Dislikes: 0
 * Total Accepted:    95.5K
 * Total Submissions: 174.3K
 * Testcase Example:  '5\n7'
 *
 * Given two integers left and right that represent the range [left, right],
 * return the bitwise AND of all numbers in this range, inclusive.
 *
 *
 * Example 1:
 *
 *
 * Input: left = 5, right = 7
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: left = 0, right = 0
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: left = 1, right = 2147483647
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= left <= right <= 2^31 - 1
 *
 *
 */

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= right - 1
	}
	return right
}

// @lc code=end

