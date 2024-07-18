/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] Course Schedule II
 *
 * https://leetcode.cn/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (58.20%)
 * Likes:    958
 * Dislikes: 0
 * Total Accepted:    238.9K
 * Total Submissions: 410.5K
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
 * Return the ordering of courses you should take to finish all courses. If
 * there are many valid answers, return any of them. If it is impossible to
 * finish all courses, return an empty array.
 *
 *
 * Example 1:
 *
 *
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: [0,1]
 * Explanation: There are a total of 2 courses to take. To take course 1 you
 * should have finished course 0. So the correct course order is [0,1].
 *
 *
 * Example 2:
 *
 *
 * Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
 * Output: [0,2,1,3]
 * Explanation: There are a total of 4 courses to take. To take course 3 you
 * should have finished both courses 1 and 2. Both courses 1 and 2 should be
 * taken after you finished course 0.
 * So one correct course order is [0,1,2,3]. Another correct ordering is
 * [0,2,1,3].
 *
 *
 * Example 3:
 *
 *
 * Input: numCourses = 1, prerequisites = []
 * Output: [0]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= numCourses * (numCourses - 1)
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * ai != bi
 * All the pairs [ai, bi] are distinct.
 *
 *
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	degree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		graph[b] = append(graph[b], a)
		degree[a]++
	}
	res := []int{}
	for i := 0; i < len(degree); i++ {
		if degree[i] == 0 {
			res = append(res, i)
		}
	}
	for i := 0; i < len(res); i++ {
		courses := graph[res[i]]
		for _, course := range courses {
			degree[course]--
			if degree[course] == 0 {
				res = append(res, course)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return []int{}
}

// @lc code=end

