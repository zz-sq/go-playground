/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.cn/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (75.75%)
 * Likes:    1797
 * Dislikes: 0
 * Total Accepted:    461.9K
 * Total Submissions: 609.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an integer array nums, return an array answer such that answer[i] is
 * equal to the product of all the elements of nums except nums[i].
 *
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 * You must write an algorithm that runs in O(n) time and without using the
 * division operation.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,4]
 * Output: [24,12,8,6]
 * Example 2:
 * Input: nums = [-1,1,0,-3,3]
 * Output: [0,0,9,0,0]
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 *
 * Follow up: Can you solve the problem in O(1) extra space complexity? (The
 * output array does not count as extra space for space complexity analysis.)
 *
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	tmp, zero := 1, 0
	for _, num := range nums {
		if num == 0 {
			zero += 1
		} else {
			tmp *= num
		}
	}
	ans := make([]int, len(nums))
	for i, num := range nums {
		if zero == 1 && num != 0 {
			ans[i] = 0
		} else if zero == 1 && num == 0 {
			ans[i] = tmp
		} else if zero > 1 {
			ans[i] = 0
		} else {
			ans[i] = tmp / num
		}
	}
	return ans
}

// @lc code=end

