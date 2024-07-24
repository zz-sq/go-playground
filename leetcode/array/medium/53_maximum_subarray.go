/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.cn/problems/maximum-subarray/description/
 *
 * algorithms
 * Medium (55.47%)
 * Likes:    6718
 * Dislikes: 0
 * Total Accepted:    1.8M
 * Total Submissions: 3.2M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the subarray with the largest sum, and
 * return its sum.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * Output: 6
 * Explanation: The subarray [4,-1,2,1] has the largest sum 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1]
 * Output: 1
 * Explanation: The subarray [1] has the largest sum 1.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [5,4,-1,7,8]
 * Output: 23
 * Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 * Follow up: If you have figured out the O(n) solution, try coding another
 * solution using the divide and conquer approach, which is more subtle.
 *
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(nums[i], nums[i]+cur)
		res = max(cur, res)
	}
	return res
}

// @lc code=end

