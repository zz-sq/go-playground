/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 *
 * https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (77.61%)
 * Likes:    872
 * Dislikes: 0
 * Total Accepted:    385.8K
 * Total Submissions: 497.1K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * Given the root of a binary search tree, and an integer k, return the k^th
 * smallest value (1-indexed) of all the values of the nodes in the tree.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,1,4,null,2], k = 1
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,3,6,2,4,null,null,1], k = 3
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is n.
 * 1 <= k <= n <= 10^4
 * 0 <= Node.val <= 10^4
 *
 *
 *
 * Follow up: If the BST is modified often (i.e., we can do insert and delete
 * operations) and you need to find the kth smallest frequently, how would you
 * optimize?
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
func kthSmallest(root *TreeNode, k int) int {
	res, pre := -1, 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || pre > k {
			return
		}
		dfs(node.Left)
		pre++
		if pre == k {
			res = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// @lc code=end

