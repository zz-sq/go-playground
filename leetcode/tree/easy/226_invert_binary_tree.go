/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] Invert Binary Tree
 *
 * https://leetcode.cn/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (80.63%)
 * Likes:    1827
 * Dislikes: 0
 * Total Accepted:    927.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * Given the root of a binary tree, invert the tree, and return its root.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [4,2,7,1,3,6,9]
 * Output: [4,7,2,9,6,3,1]
 *
 *
 * Example 2:
 *
 *
 * Input: root = [2,1,3]
 * Output: [2,3,1]
 *
 *
 * Example 3:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	head := &TreeNode{}
	head.Val = root.Val
	head.Left = invertTree(root.Right)
	head.Right = invertTree(root.Left)
	return head
}

// @lc code=end

