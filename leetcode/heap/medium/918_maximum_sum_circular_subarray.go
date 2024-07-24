/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] Maximum Sum Circular Subarray
 *
 * https://leetcode.cn/problems/maximum-sum-circular-subarray/description/
 *
 * algorithms
 * Medium (42.40%)
 * Likes:    714
 * Dislikes: 0
 * Total Accepted:    92.3K
 * Total Submissions: 217.7K
 * Testcase Example:  '[1,-2,3,-2]'
 *
 * Given a circular integer array nums of length n, return the maximum possible
 * sum of a non-empty subarray of nums.
 *
 * A circular array means the end of the array connects to the beginning of the
 * array. Formally, the next element of nums[i] is nums[(i + 1) % n] and the
 * previous element of nums[i] is nums[(i - 1 + n) % n].
 *
 * A subarray may only include each element of the fixed buffer nums at most
 * once. Formally, for a subarray nums[i], nums[i + 1], ..., nums[j], there
 * does not exist i <= k1, k2 <= j with k1 % n == k2 % n.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,-2,3,-2]
 * Output: 3
 * Explanation: Subarray [3] has maximum sum 3.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [5,-3,5]
 * Output: 10
 * Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [-3,-2,-3]
 * Output: -2
 * Explanation: Subarray [-2] has maximum sum -2.
 *
 *
 *
 * Constraints:
 *
 *
 * n == nums.length
 * 1 <= n <= 3 * 10^4
 * -3 * 10^4 <= nums[i] <= 3 * 10^4
 *
 *
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxRes, maxCur, sum := nums[0], nums[0], nums[0]
	minRes, minCur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxCur = max(nums[i], maxCur+nums[i])
		maxRes = max(maxCur, maxRes)
		minCur = min(nums[i], minCur+nums[i])
		minRes = min(minCur, minRes)
		sum += nums[i]
	}
	if sum == minRes {
		return maxRes
	}
	return max(maxRes, sum-minRes)
}

// @lc code=end

