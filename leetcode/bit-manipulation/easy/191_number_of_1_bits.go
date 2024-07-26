/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] Number of 1 Bits
 *
 * https://leetcode.cn/problems/number-of-1-bits/description/
 *
 * algorithms
 * Easy (77.69%)
 * Likes:    643
 * Dislikes: 0
 * Total Accepted:    388.3K
 * Total Submissions: 499.8K
 * Testcase Example:  '11'
 *
 * Write a function that takes the binary representation of a positive integer
 * and returns the number of set bits it has (also known as the Hamming
 * weight).
 *
 *
 * Example 1:
 *
 *
 * Input: n = 11
 *
 * Output: 3
 *
 * Explanation:
 *
 * The input binary string 1011 has a total of three set bits.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 128
 *
 * Output: 1
 *
 * Explanation:
 *
 * The input binary string 10000000 has a total of one set bit.
 *
 *
 * Example 3:
 *
 *
 * Input: n = 2147483645
 *
 * Output: 30
 *
 * Explanation:
 *
 * The input binary string 1111111111111111111111111111101 has a total of
 * thirty set bits.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 2^31 - 1
 *
 *
 *
 * Follow up: If this function is called many times, how would you optimize it?
 */

// @lc code=start
func hammingWeight(n int) int {
	res := 0
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}

// @lc code=end

