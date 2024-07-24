/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (42.42%)
 * Likes:    7174
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 2.7M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return
 * the median of the two sorted arrays.
 *
 * The overall run time complexity should be O(log (m+n)).
 *
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,3], nums2 = [2]
 * Output: 2.00000
 * Explanation: merged array = [1,2,3] and median is 2.
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [1,2], nums2 = [3,4]
 * Output: 2.50000
 * Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 *
 *
 *
 * Constraints:
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	mid := (m + n + 1) / 2

	if (m+n)%2 == 0 {
		return (float64(getKth(nums1, 0, m-1, nums2, 0, n-1, mid)) +
			float64(getKth(nums1, 0, m-1, nums2, 0, n-1, mid+1))) / 2.0
	} else {
		return float64(getKth(nums1, 0, m-1, nums2, 0, n-1, mid))
	}
}

func getKth(nums1 []int, s1 int, e1 int, nums2 []int, s2 int, e2 int, k int) int {
	l1 := e1 - s1 + 1
	l2 := e2 - s2 + 1

	if l1 == 0 {
		return nums2[s2+k-1]
	} else if l2 == 0 {
		return nums1[s1+k-1]
	}

	if k == 1 {
		return min(nums1[s1], nums2[s2])
	}

	i := s1 + min(k/2, l1) - 1
	j := s2 + min(k/2, l2) - 1

	if nums1[i] < nums2[j] {
		return getKth(nums1, i+1, e1, nums2, s2, e2, k-(i-s1+1))
	} else {
		return getKth(nums1, s1, e1, nums2, j+1, e2, k-(j-s2+1))
	}
}

// @lc code=end

