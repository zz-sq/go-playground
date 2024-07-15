/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.cn/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (38.11%)
 * Likes:    2373
 * Dislikes: 0
 * Total Accepted:    944K
 * Total Submissions: 2.5M
 * Testcase Example:  '[2,1,3]'
 *
 * Given the root of a binary tree, determine if it is a valid binary search
 * tree (BST).
 *
 * A valid BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [2,1,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is
 * 4.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * -2^31 <= Node.val <= 2^31 - 1
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
func isValidBST(root *TreeNode) bool {
	res := true
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, maxNum int, minNum int) {
		if node == nil || !res {
			return
		}
		if node.Val >= maxNum || node.Val <= minNum {
			res = false
			return
		}
		dfs(node.Left, min(node.Val, maxNum), minNum)
		dfs(node.Right, maxNum, max(node.Val, minNum))
	}
	dfs(root, math.MaxInt64, math.MinInt64)
	return res
}

// @lc code=end

