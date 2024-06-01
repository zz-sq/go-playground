/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
 *
 * https://leetcode.cn/problems/leaf-similar-trees/description/
 *
 * algorithms
 * Easy (65.14%)
 * Likes:    235
 * Dislikes: 0
 * Total Accepted:    88.7K
 * Total Submissions: 136.1K
 * Testcase Example:  '[3,5,1,6,2,9,8,null,null,7,4]\n' +
  '[3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]'
 *
 * Consider all the leaves of a binary tree, from left to right order, the
 * values of those leaves form a leaf value sequence.
 *
 *
 *
 * For example, in the given tree above, the leaf value sequence is (6, 7, 4,
 * 9, 8).
 *
 * Two binary trees are considered leaf-similar if their leaf value sequence is
 * the same.
 *
 * Return true if and only if the two given trees with head nodes root1 and
 * root2 are leaf-similar.
 *
 *
 * Example 1:
 *
 *
 * Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 =
 * [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root1 = [1,2,3], root2 = [1,3,2]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in each tree will be in the range [1, 200].
 * Both of the given trees will have values in the range [0, 200].
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

func dfs(root *TreeNode, l *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*l = append(*l, root.Val)
	}

	dfs(root.Left, l)
	dfs(root.Right, l)
}

func stack(root *TreeNode) (ret []int) {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	next := root
	for {
		for next != nil {
			if next.Left == nil && next.Right == nil {
				ret = append(ret, next.Val)
			}
			if next.Right != nil {
				stack = append(stack, next.Right)
			}
			next = next.Left
		}
		if len(stack) == 0 {
			break
		}
		next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// l1, l2 := []int{}, []int{}
	// dfs(root1, &l1)
	// dfs(root2, &l2)

	l1 := stack(root1)
	l2 := stack(root2)

	if len(l1) != len(l2) {
		return false
	}

	for i, v := range l1 {
		if v != l2[i] {
			return false
		}
	}
	return true
}

// @lc code=end

