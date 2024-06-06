/*
 * @lc app=leetcode.cn id=2769 lang=golang
 *
 * [2769] Find the Maximum Achievable Number
 *
 * https://leetcode.cn/problems/find-the-maximum-achievable-number/description/
 *
 * algorithms
 * Easy (92.78%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    36.9K
 * Total Submissions: 39.8K
 * Testcase Example:  '4\n1'
 *
 * Given two integers, num and t. A number is achievable if it can become equal
 * to num after applying the following operation:
 *
 *
 * Increase or decrease the number by 1, and simultaneously increase or
 * decrease num by 1.
 *
 *
 * Return the maximum achievable number after applying the operation at most t
 * times.
 *
 *
 * Example 1:
 *
 *
 * Input: num = 4, t = 1
 *
 * Output: 6
 *
 * Explanation:
 *
 * Apply the following operation once to make the maximum achievable number
 * equal to num:
 *
 *
 * Decrease the maximum achievable number by 1, and increase num by 1.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: num = 3, t = 2
 *
 * Output: 7
 *
 * Explanation:
 *
 * Apply the following operation twice to make the maximum achievable number
 * equal to num:
 *
 *
 * Decrease the maximum achievable number by 1, and increase num by 1.
 *
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= num, t <= 50
 *
 *
 */

// @lc code=start
func theMaximumAchievableX(num int, t int) int {
	return num + 2*t
}

// @lc code=end

