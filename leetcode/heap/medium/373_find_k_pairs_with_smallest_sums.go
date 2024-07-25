/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 *
 * https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
 *
 * algorithms
 * Medium (41.06%)
 * Likes:    604
 * Dislikes: 0
 * Total Accepted:    81.2K
 * Total Submissions: 197.6K
 * Testcase Example:  '[1,7,11]\n[2,4,6]\n3'
 *
 * You are given two integer arrays nums1 and nums2 sorted in non-decreasing
 * order and an integer k.
 *
 * Define a pair (u, v) which consists of one element from the first array and
 * one element from the second array.
 *
 * Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest
 * sums.
 *
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
 * Output: [[1,2],[1,4],[1,6]]
 * Explanation: The first 3 pairs are returned from the sequence:
 * [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
 * Output: [[1,1],[1,1]]
 * Explanation: The first 2 pairs are returned from the sequence:
 * [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums1.length, nums2.length <= 10^5
 * -10^9 <= nums1[i], nums2[i] <= 10^9
 * nums1 and nums2 both are sorted in non-decreasing order.
 * 1 <= k <= 10^4
 * k <= nums1.length * nums2.length
 *
 *
 */

// @lc code=start
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	res := make([][]int, 0, min(k, n*m))
	h := make(hp, min(k, n))
	for i := range h {
		h[i] = tuple{nums1[i] + nums2[0], i, 0}
	}
	for len(res) < k {
		p := heap.Pop(&h).(tuple)
		i, j := p.i, p.j
		res = append(res, []int{nums1[i], nums2[j]})
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return res
}

type tuple struct{ s, i, j int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].s < h[j].s }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// @lc code=end

