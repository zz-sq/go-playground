/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (61.25%)
 * Likes:    2504
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.8M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Given an integer array nums and an integer k, return the k^th largest
 * element in the array.
 *
 * Note that it is the k^th largest element in the sorted order, not the k^th
 * distinct element.
 *
 * Can you solve it without sorting?
 *
 *
 * Example 1:
 * Input: nums = [3,2,1,5,6,4], k = 2
 * Output: 5
 * Example 2:
 * Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
 * Output: 4
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		maxHeapify(nums, i, n)
	}
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		n--
		maxHeapify(nums, 0, n)
	}
	return nums[0]
}

func maxHeapify(a []int, i, n int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < n && a[l] > a[largest] {
		largest = l
	}
	if r < n && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, n)
	}
}

// @lc code=end

