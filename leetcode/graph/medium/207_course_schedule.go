/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] Course Schedule
 *
 * https://leetcode.cn/problems/course-schedule/description/
 *
 * algorithms
 * Medium (54.16%)
 * Likes:    1976
 * Dislikes: 0
 * Total Accepted:    427.3K
 * Total Submissions: 789K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of numCourses courses you have to take, labeled from 0 to
 * numCourses - 1. You are given an array prerequisites where prerequisites[i]
 * = [ai, bi] indicates that you must take course bi first if you want to take
 * course ai.
 *
 *
 * For example, the pair [0, 1], indicates that to take course 0 you have to
 * first take course 1.
 *
 *
 * Return true if you can finish all courses. Otherwise, return false.
 *
 *
 * Example 1:
 *
 *
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0. So it is possible.
 *
 *
 * Example 2:
 *
 *
 * Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0, and to take course 0 you
 * should also have finished course 1. So it is impossible.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= 5000
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * All the pairs prerequisites[i] are unique.
 *
 *
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	maps := map[int]map[int]bool{}
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		if a == b {
			return false
		}
		if dfs(maps, a, b, map[int]bool{}) {
			return false
		}
		if _, ok := maps[b]; !ok {
			maps[b] = map[int]bool{}
		}
		maps[b][a] = true
	}
	return true
}

func dfs(maps map[int]map[int]bool, a int, b int, visited map[int]bool) bool {
	mapA, ok := maps[a]
	if !ok {
		return false
	}
	if mapA[b] {
		return true
	}
	for c := range mapA {
		if visited[c] {
			continue
		}
		if dfs(maps, c, b, visited) {
			return true
		}
		visited[c] = true
	}
	return false
}

// @lc code=end

