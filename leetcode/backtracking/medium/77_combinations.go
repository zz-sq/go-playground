/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.cn/problems/combinations/description/
 *
 * algorithms
 * Medium (77.08%)
 * Likes:    1646
 * Dislikes: 0
 * Total Accepted:    739.7K
 * Total Submissions: 959.6K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * chosen from the range [1, n].
 *
 * You may return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 4, k = 2
 * Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
 * Explanation: There are 4 choose 2 = 6 total combinations.
 * Note that combinations are unordered, i.e., [1,2] and [2,1] are considered
 * to be the same combination.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1, k = 1
 * Output: [[1]]
 * Explanation: There is 1 choose 1 = 1 total combination.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 20
 * 1 <= k <= n
 *
 *
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	if k == 0 {
		return res
	}
	var backstrace func(j int, cur []int, visited []bool)
	backstrace = func(j int, cur []int, visited []bool) {
		if k == len(cur) {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := j; i <= n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			cur = append(cur, i)
			backstrace(i+1, cur, visited)
			visited[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	visited := make([]bool, n+1)
	backstrace(1, []int{}, visited)
	return res
}

// @lc code=end

