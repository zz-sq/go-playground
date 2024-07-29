/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.cn/problems/triangle/description/
 *
 * algorithms
 * Medium (68.82%)
 * Likes:    1356
 * Dislikes: 0
 * Total Accepted:    363.5K
 * Total Submissions: 528.2K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle array, return the minimum path sum from top to bottom.
 *
 * For each step, you may move to an adjacent number of the row below. More
 * formally, if you are on index i on the current row, you may move to either
 * index i or index i + 1 on the next row.
 *
 *
 * Example 1:
 *
 *
 * Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
 * Output: 11
 * Explanation: The triangle looks like:
 * ⁠  2
 * ⁠ 3 4
 * ⁠6 5 7
 * 4 1 8 3
 * The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined
 * above).
 *
 *
 * Example 2:
 *
 *
 * Input: triangle = [[-10]]
 * Output: -10
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= triangle.length <= 200
 * triangle[0].length == 1
 * triangle[i].length == triangle[i - 1].length + 1
 * -10^4 <= triangle[i][j] <= 10^4
 *
 *
 *
 * Follow up: Could you do this using only O(n) extra space, where n is the
 * total number of rows in the triangle?
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = math.MaxInt
	}
	res := math.MaxInt
	for i := 0; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if i == 0 && j == 0 {
				dp[j] = triangle[i][j]
			} else if j == 0 {
				dp[j] = dp[j] + triangle[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
			}
			if i == len(triangle)-1 {
				res = min(res, dp[j])
			}
		}
	}
	return res
}

// @lc code=end

