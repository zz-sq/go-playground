/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] Rotate List
 *
 * https://leetcode.cn/problems/rotate-list/description/
 *
 * algorithms
 * Medium (41.38%)
 * Likes:    1065
 * Dislikes: 0
 * Total Accepted:    393.9K
 * Total Submissions: 951.8K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, rotate the list to the right by k
 * places.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], k = 2
 * Output: [4,5,1,2,3]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [0,1,2], k = 4
 * Output: [2,0,1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 500].
 * -100 <= Node.val <= 100
 * 0 <= k <= 2 * 10^9
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	pre, cur := head, head
	for i := 0; i < k; i++ {
		if cur.Next == nil {
			return rotateRight(head, k%(i+1))
		}
		cur = cur.Next
	}
	for cur.Next != nil {
		cur, pre = cur.Next, pre.Next
	}
	cur.Next = head
	newHead := pre.Next
	pre.Next = nil
	return newHead
}

// @lc code=end

