/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.cn/problems/add-binary/description/
 *
 * algorithms
 * Easy (53.22%)
 * Likes:    1214
 * Dislikes: 0
 * Total Accepted:    408.8K
 * Total Submissions: 768.1K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings a and b, return their sum as a binary string.
 *
 *
 * Example 1:
 * Input: a = "11", b = "1"
 * Output: "100"
 * Example 2:
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 *
 * Constraints:
 *
 *
 * 1 <= a.length, b.length <= 10^4
 * a and b consist only of '0' or '1' characters.
 * Each string does not contain leading zeros except for the zero itself.
 *
 *
 */

// @lc code=start
func addBinary(a string, b string) string {
	res := []byte{}
	i, j, flag := len(a)-1, len(b)-1, byte(0)

	for i >= 0 || j >= 0 || flag > 0 {
		if i >= 0 {
			flag += a[i] - '0'
			i--
		}
		if j >= 0 {
			flag += b[j] - '0'
			j--
		}
		res = append([]byte{flag%2 + '0'}, res...)
		flag /= 2
	}
	return string(res)
}

// @lc code=end

