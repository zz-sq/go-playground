package main

import (
	"fmt"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func minSubArrayLen(target int, nums []int) int {
	left, sum, res := 0, 0, len(nums)+1
	for i, n := range nums {
		sum += n
		for sum >= target {
			sum -= nums[left]
			if res > i-left+1 {
				res = i - left + 1
			}
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	} else {
		return res
	}
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{7, 3, 1, 2, 4, 3}))
}
