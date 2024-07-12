/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 *
 * https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (73.90%)
 * Likes:    1693
 * Dislikes: 0
 * Total Accepted:    488.9K
 * Total Submissions: 661.5K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * Given the root of a binary tree, flatten the tree into a "linked
 * list":
 *
 *
 * The "linked list" should use the same TreeNode class where the right child
 * pointer points to the next node in the list and the left child pointer is
 * always null.
 * The "linked list" should be in the same order as a pre-order traversal of
 * the binary tree.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,5,3,4,null,6]
 * Output: [1,null,2,null,3,null,4,null,5,null,6]
 *
 *
 * Example 2:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: root = [0]
 * Output: [0]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 2000].
 * -100 <= Node.val <= 100
 *
 *
 *
 * Follow up: Can you flatten the tree in-place (with O(1) extra space)?
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

func flattenDfs(root *TreeNode) {
	if root == nil {
		return
	}
	res := dfs(root)

	for i := 1; i < len(res); i++ {
		pre, cur := res[i-1], res[i]
		pre.Left, pre.Right = nil, cur
	}
}

func dfs(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	if root != nil {
		res = append(res, root)
		res = append(res, dfs(root.Left)...)
		res = append(res, dfs(root.Right)...)
	}
	return res
}

func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			left := root.Left
			for left.Right != nil {
				left = left.Right
			}
			left.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}

// @lc code=end

