/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 *
 * https://leetcode.cn/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (67.32%)
 * Likes:    1086
 * Dislikes: 0
 * Total Accepted:    440.9K
 * Total Submissions: 654.9K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * Given the root of a binary tree, imagine yourself standing on the right side
 * of it, return the values of the nodes you can see ordered from top to
 * bottom.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3,null,5,null,4]
 * Output: [1,3,4]
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,null,3]
 * Output: [1,3]
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
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		next := []*TreeNode{}
		curLevel := []int{}
		for _, v := range nodes {
			curLevel = append(curLevel, v.Val)
			if v.Left != nil {
				next = append(next, v.Left)
			}
			if v.Right != nil {
				next = append(next, v.Right)
			}
		}
		res = append(res, curLevel[len(curLevel)-1])
		nodes = next
	}
	return res
}

// @lc code=end

