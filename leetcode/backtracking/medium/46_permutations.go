/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.cn/problems/permutations/description/
 *
 * algorithms
 * Medium (79.38%)
 * Likes:    2923
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3]'
 *
 * Given an array nums of distinct integers, return all the possible
 * permutations. You can return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * Example 2:
 * Input: nums = [0,1]
 * Output: [[0,1],[1,0]]
 * Example 3:
 * Input: nums = [1]
 * Output: [[1]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * All the integers of nums are unique.
 *
 *
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	var backstrace func(cur []int, visited []bool)
	backstrace = func(cur []int, visited []bool) {
		if len(nums) == len(cur) {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			cur = append(cur, nums[i])
			backstrace(cur, visited)
			visited[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	visited := make([]bool, len(nums))
	backstrace([]int{}, visited)
	return res
}

// @lc code=end

