/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.cn/problems/powx-n/description/
 *
 * algorithms
 * Medium (38.39%)
 * Likes:    1363
 * Dislikes: 0
 * Total Accepted:    463.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (i.e.,
 * x^n).
 *
 *
 * Example 1:
 *
 *
 * Input: x = 2.00000, n = 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: x = 2.10000, n = 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: x = 2.00000, n = -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 *
 * Constraints:
 *
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * n is an integer.
 * Either x is not zero or n > 0.
 * -10^4 <= x^n <= 10^4
 *
 *
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	res, y := 1.0, x
	for n > 0 {
		if n%2 == 1 {
			res *= y
		}
		y *= y
		n /= 2
	}
	return res
}

// @lc code=end

