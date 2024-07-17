/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 *
 * https://leetcode.cn/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Easy (81.78%)
 * Likes:    1149
 * Dislikes: 0
 * Total Accepted:    411.7K
 * Total Submissions: 503.4K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * Given the root of a complete binary tree, return the number of the nodes in
 * the tree.
 *
 * According to Wikipedia, every level, except possibly the last, is completely
 * filled in a complete binary tree, and all nodes in the last level are as far
 * left as possible. It can have between 1 and 2^h nodes inclusive at the last
 * level h.
 *
 * Design an algorithm that runs in less than O(n) time complexity.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3,4,5,6]
 * Output: 6
 *
 *
 * Example 2:
 *
 *
 * Input: root = []
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 5 * 10^4].
 * 0 <= Node.val <= 5 * 10^4
 * The tree is guaranteed to be complete.
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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxHeight(root.Left)
	right := maxHeight(root.Right)
	if left == right {
		return countNodes(root.Right) + (1 << left)
	}
	return countNodes(root.Left) + (1 << right)
}

func maxHeight(root *TreeNode) int {
	height := 0
	for root != nil {
		height++
		root = root.Left
	}
	return height
}

// @lc code=end

