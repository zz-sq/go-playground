package main

import (
	"fmt"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
// https://leetcode.cn/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150
func removeElement(nums []int, val int) int {
	i, j, k := 0, 0, 0
	for ; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		} else {
			k++
		}
	}
	return k
}

func main() {
	nums1 := []int{1, 2, 2, 0, 0, 0}
	fmt.Println(removeElement(nums1, 2))
	fmt.Println(nums1)
}
