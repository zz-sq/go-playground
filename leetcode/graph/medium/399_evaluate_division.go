/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] Evaluate Division
 *
 * https://leetcode.cn/problems/evaluate-division/description/
 *
 * algorithms
 * Medium (58.63%)
 * Likes:    1108
 * Dislikes: 0
 * Total Accepted:    101.5K
 * Total Submissions: 173.1K
 * Testcase Example:  '[["a","b"],["b","c"]]\n' +
  '[2.0,3.0]\n' +
  '[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]'
 *
 * You are given an array of variable pairs equations and an array of real
 * numbers values, where equations[i] = [Ai, Bi] and values[i] represent the
 * equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a
 * single variable.
 *
 * You are also given some queries, where queries[j] = [Cj, Dj] represents the
 * j^th query where you must find the answer for Cj / Dj = ?.
 *
 * Return the answers to all queries. If a single answer cannot be determined,
 * return -1.0.
 *
 * Note: The input is always valid. You may assume that evaluating the queries
 * will not result in division by zero and that there is no contradiction.
 *
 * Note: The variables that do not occur in the list of equations are
 * undefined, so the answer cannot be determined for them.
 *
 *
 * Example 1:
 *
 *
 * Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries =
 * [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
 * Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
 * Explanation:
 * Given: a / b = 2.0, b / c = 3.0
 * queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
 * return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
 * note: x is undefined => -1.0
 *
 * Example 2:
 *
 *
 * Input: equations = [["a","b"],["b","c"],["bc","cd"]], values =
 * [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
 * Output: [3.75000,0.40000,5.00000,0.20000]
 *
 *
 * Example 3:
 *
 *
 * Input: equations = [["a","b"]], values = [0.5], queries =
 * [["a","b"],["b","a"],["a","c"],["x","y"]]
 * Output: [0.50000,2.00000,-1.00000,-1.00000]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= equations.length <= 20
 * equations[i].length == 2
 * 1 <= Ai.length, Bi.length <= 5
 * values.length == equations.length
 * 0.0 < values[i] <= 20.0
 * 1 <= queries.length <= 20
 * queries[i].length == 2
 * 1 <= Cj.length, Dj.length <= 5
 * Ai, Bi, Cj, Dj consist of lower case English letters and digits.
 *
 *
*/

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	vals := map[string]map[string]float64{}
	for i := 0; i < len(equations); i++ {
		a, b := equations[i][0], equations[i][1]
		if _, ok := vals[a]; !ok {
			vals[a] = map[string]float64{}
		}
		vals[a][b] = values[i]
		if _, ok := vals[b]; !ok {
			vals[b] = map[string]float64{}
		}
		vals[b][a] = 1.0 / values[i]
	}
	res := []float64{}
	for i := 0; i < len(queries); i++ {
		res = append(res, dfs(vals, queries[i][0], queries[i][1], map[string]bool{}))
	}
	return res
}

func dfs(vals map[string]map[string]float64, a string, b string, visited map[string]bool) float64 {
	val, ok := vals[a]
	if !ok {
		return -1.0
	}
	v, ok := val[b]
	if ok {
		return v
	}
	visited[a] = true
	for c, n := range val {
		if visited[c] {
			continue
		}
		next := dfs(vals, c, b, visited)
		if next != -1.0 {
			return n * next
		}
	}
	return -1.0
}

// @lc code=end

