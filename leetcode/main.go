package main

import (
	"slices"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
// https://leetcode.cn/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150
func rotate(nums []int, k int) {
	k %= len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums1, 3)
}
