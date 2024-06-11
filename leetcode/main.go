package main

import (
	"fmt"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
// https://leetcode.cn/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150
func majorityElement(nums []int) int {
	votes, res := 0, 0
	for _, v := range nums {
		if votes == 0 {
			res = v
		}
		if res == v {
			votes++
		} else {
			votes--
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 1, 1, 1, 2, 3}
	fmt.Println(majorityElement(nums1))
	fmt.Println(nums1)
}
