/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 *
 * https://leetcode.cn/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (46.48%)
 * Likes:    2159
 * Dislikes: 0
 * Total Accepted:    795.2K
 * Total Submissions: 1.7M
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * Given an array of positive integers nums and a positive integer target,
 * return the minimal length of a subarray whose sum is greater than or equal
 * to target. If there is no such subarray, return 0 instead.
 *
 *
 * Example 1:
 *
 *
 * Input: target = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: The subarray [4,3] has the minimal length under the problem
 * constraint.
 *
 *
 * Example 2:
 *
 *
 * Input: target = 4, nums = [1,4,4]
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: target = 11, nums = [1,1,1,1,1,1,1,1]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= target <= 10^9
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^4
 *
 *
 *
 * Follow up: If you have figured out the O(n) solution, try coding another
 * solution of which the time complexity is O(n log(n)).
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	left, sum, res := 0, 0, len(nums)+1
	for i, n := range nums {
		sum += n
		for sum >= target {
			sum -= nums[left]
			if res > i-left+1 {
				res = i - left + 1
			}
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	} else {
		return res
	}
}

// @lc code=end

