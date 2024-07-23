/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.cn/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.87%)
 * Likes:    2841
 * Dislikes: 0
 * Total Accepted:    981.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given an array of distinct integers candidates and a target integer target,
 * return a list of all unique combinations of candidates where the chosen
 * numbers sum to target. You may return the combinations in any order.
 *
 * The same number may be chosen from candidates an unlimited number of times.
 * Two combinations are unique if the frequency of at least one of the chosen
 * numbers is different.
 *
 * The test cases are generated such that the number of unique combinations
 * that sum up to target is less than 150 combinations for the given input.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7
 * Output: [[2,2,3],[7]]
 * Explanation:
 * 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple
 * times.
 * 7 is a candidate, and 7 = 7.
 * These are the only two combinations.
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8
 * Output: [[2,2,2,2],[2,3,3],[3,5]]
 *
 *
 * Example 3:
 *
 *
 * Input: candidates = [2], target = 1
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= candidates.length <= 30
 * 2 <= candidates[i] <= 40
 * All elements of candidates are distinct.
 * 1 <= target <= 40
 *
 *
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	var dfs func(target int, start int, cur []int)
	dfs = func(target int, start int, cur []int) {
		if target == 0 {
			tmp := []int{}
			tmp = append(tmp, cur...)
			res = append(res, tmp)
		}
		for i := start; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break
			}
			cur = append(cur, candidates[i])
			dfs(target-candidates[i], i, cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target, 0, []int{})
	return res
}

// @lc code=end

