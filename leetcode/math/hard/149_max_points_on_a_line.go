/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] Max Points on a Line
 *
 * https://leetcode.cn/problems/max-points-on-a-line/description/
 *
 * algorithms
 * Hard (40.77%)
 * Likes:    561
 * Dislikes: 0
 * Total Accepted:    96.5K
 * Total Submissions: 236.6K
 * Testcase Example:  '[[1,1],[2,2],[3,3]]'
 *
 * Given an array of points where points[i] = [xi, yi] represents a point on
 * the X-Y plane, return the maximum number of points that lie on the same
 * straight line.
 *
 *
 * Example 1:
 *
 *
 * Input: points = [[1,1],[2,2],[3,3]]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
 * Output: 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= points.length <= 300
 * points[i].length == 2
 * -10^4 <= xi, yi <= 10^4
 * All the points are unique.
 *
 *
 */

// @lc code=start
func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}
	res := 0
	for i := 0; i < len(points); i++ {
		cur, duplicate := 0, 0
		m := make(map[string]int)
		for j := i + 1; j < len(points); j++ {
			x := points[j][0] - points[i][0]
			y := points[j][1] - points[i][1]
			if x == 0 && y == 0 {
				duplicate++
				continue
			}
			gcdVal := gcd(x, y)
			x, y = x/gcdVal, y/gcdVal
			key := fmt.Sprintf("%d %d", x, y)
			m[key]++
			cur = max(cur, m[key])
		}
		res = max(res, cur+duplicate+1)
	}
	return res
}

func gcd(x int, y int) int {
	for y != 0 {
		t := x % y
		x, y = y, t
	}
	return x
}

// @lc code=end

