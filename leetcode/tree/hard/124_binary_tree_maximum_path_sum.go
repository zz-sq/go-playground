/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 *
 * https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (45.66%)
 * Likes:    2247
 * Dislikes: 0
 * Total Accepted:    428.9K
 * Total Submissions: 939.3K
 * Testcase Example:  '[1,2,3]'
 *
 * A path in a binary tree is a sequence of nodes where each pair of adjacent
 * nodes in the sequence has an edge connecting them. A node can only appear in
 * the sequence at most once. Note that the path does not need to pass through
 * the root.
 *
 * The path sum of a path is the sum of the node's values in the path.
 *
 * Given the root of a binary tree, return the maximum path sum of any
 * non-empty path.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3]
 * Output: 6
 * Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 =
 * 6.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [-10,9,20,null,null,15,7]
 * Output: 42
 * Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 +
 * 7 = 42.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 3 * 10^4].
 * -1000 <= Node.val <= 1000
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := max(dfs(node.Left), 0)
		r := max(dfs(node.Right), 0)
		res = max(res, l+r+node.Val)
		return node.Val + max(l, r)
	}
	dfs(root)
	return res
}

// @lc code=end

