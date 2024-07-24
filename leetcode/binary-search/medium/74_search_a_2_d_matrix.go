/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.cn/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (49.92%)
 * Likes:    942
 * Dislikes: 0
 * Total Accepted:    428K
 * Total Submissions: 857.3K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * You are given an m x n integer matrix matrix with the following two
 * properties:
 *
 *
 * Each row is sorted in non-decreasing order.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 * Given an integer target, return true if target is in matrix or false
 * otherwise.
 *
 * You must write a solution in O(log(m * n)) time complexity.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 100
 * -10^4 <= matrix[i][j], target <= 10^4
 *
 *
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	l, r := 0, len(matrix)*len(matrix[0])-1
	for l <= r {
		mid := l + (r-l)/2
		i, j := mid/len(matrix[0]), mid%len(matrix[0])
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

// @lc code=end

