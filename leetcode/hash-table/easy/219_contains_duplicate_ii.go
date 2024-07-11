/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] Contains Duplicate II
 *
 * https://leetcode.cn/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (46.95%)
 * Likes:    712
 * Dislikes: 0
 * Total Accepted:    313K
 * Total Submissions: 666.7K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * Given an integer array nums and an integer k, return true if there are two
 * distinct indices i and j in the array such that nums[i] == nums[j] and abs(i
 * - j) <= k.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1], k = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,0,1,1], k = 1
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,3,1,2,3], k = 2
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 0 <= k <= 10^5
 *
 *
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	cnt := map[int]int{}
	for i, num := range nums {
		if j, ok := cnt[num]; ok {
			if i-j <= k {
				return true
			}
		}
		cnt[num] = i
	}
	return false
}

// @lc code=end

