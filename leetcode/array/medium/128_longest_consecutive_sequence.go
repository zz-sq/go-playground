/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.cn/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (51.63%)
 * Likes:    2130
 * Dislikes: 0
 * Total Accepted:    688K
 * Total Submissions: 1.3M
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers nums, return the length of the longest
 * consecutive elements sequence.
 *
 * You must write an algorithm that runs in O(n) time.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [100,4,200,1,3,2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,3,7,2,5,8,4,6,0,1]
 * Output: 9
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	cnt := map[int]bool{}
	for _, num := range nums {
		cnt[num] = true
	}
	res := 0
	for num := range cnt {
		if cnt[num-1] {
			continue
		}
		cur := num
		for cnt[cur] {
			cur++
		}
		res = max(res, cur-num)
	}
	return res
}

// @lc code=end

