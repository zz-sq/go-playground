/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] Partition List
 *
 * https://leetcode.cn/problems/partition-list/description/
 *
 * algorithms
 * Medium (64.84%)
 * Likes:    845
 * Dislikes: 0
 * Total Accepted:    291.9K
 * Total Submissions: 450.2K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given the head of a linked list and a value x, partition it such that all
 * nodes less than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,4,3,2,5,2], x = 3
 * Output: [1,2,2,4,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [2,1], x = 2
 * Output: [1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 200].
 * -100 <= Node.val <= 100
 * -200 <= x <= 200
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	low, high := &ListNode{}, &ListNode{}
	curLow, curHigh := low, high
	for head != nil {
		if head.Val < x {
			curLow.Next = head
			curLow = curLow.Next
		} else {
			curHigh.Next = head
			curHigh = curHigh.Next
		}
		head = head.Next
	}
	curLow.Next = high.Next
	curHigh.Next = nil
	return low.Next
}

// @lc code=end

