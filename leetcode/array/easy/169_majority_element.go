/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.cn/problems/majority-element/description/
 *
 * algorithms
 * Easy (66.37%)
 * Likes:    2214
 * Dislikes: 0
 * Total Accepted:    971.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array nums of size n, return the majority element.
 *
 * The majority element is the element that appears more than ⌊n / 2⌋ times.
 * You may assume that the majority element always exists in the array.
 *
 *
 * Example 1:
 * Input: nums = [3,2,3]
 * Output: 3
 * Example 2:
 * Input: nums = [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 * Constraints:
 *
 *
 * n == nums.length
 * 1 <= n <= 5 * 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 *
 * Follow-up: Could you solve the problem in linear time and in O(1) space?
 */

// @lc code=start
func majorityElement(nums []int) int {
	votes, res := 0, 0
	for _, v := range nums {
		if votes == 0 {
			res = v
		}
		if res == v {
			votes++
		} else {
			votes--
		}
	}
	return res
}

// @lc code=end

