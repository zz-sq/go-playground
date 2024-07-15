/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] Summary Ranges
 *
 * https://leetcode.cn/problems/summary-ranges/description/
 *
 * algorithms
 * Easy (54.76%)
 * Likes:    391
 * Dislikes: 0
 * Total Accepted:    170.2K
 * Total Submissions: 310.7K
 * Testcase Example:  '[0,1,2,4,5,7]'
 *
 * You are given a sorted unique integer array nums.
 *
 * A range [a,b] is the set of all integers from a to b (inclusive).
 *
 * Return the smallest sorted list of ranges that cover all the numbers in the
 * array exactly. That is, each element of nums is covered by exactly one of
 * the ranges, and there is no integer x such that x is in one of the ranges
 * but not in nums.
 *
 * Each range [a,b] in the list should be output as:
 *
 *
 * "a->b" if a != b
 * "a" if a == b
 *
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [0,1,2,4,5,7]
 * Output: ["0->2","4->5","7"]
 * Explanation: The ranges are:
 * [0,2] --> "0->2"
 * [4,5] --> "4->5"
 * [7,7] --> "7"
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,2,3,4,6,8,9]
 * Output: ["0","2->4","6","8->9"]
 * Explanation: The ranges are:
 * [0,0] --> "0"
 * [2,4] --> "2->4"
 * [6,6] --> "6"
 * [8,9] --> "8->9"
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 20
 * -2^31 <= nums[i] <= 2^31 - 1
 * All the values of nums are unique.
 * nums is sorted in ascending order.
 *
 *
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	res := []string{}
	start := 0
	for i := range nums {
		if i == len(nums)-1 || nums[i+1]-1 != nums[i] {
			if i == start {
				res = append(res, fmt.Sprintf("%d", nums[i]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
			}
			start = i + 1
		}
	}
	return res
}

// @lc code=end

