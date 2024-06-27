package main

import (
	"fmt"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func maxArea(height []int) int {
	res, l, r := 0, 0, len(height)-1
	for l < r {
		minHeight := height[l]
		if height[r] < minHeight {
			minHeight = height[r]
		}
		cur := (r - l) * minHeight
		if res < cur {
			res = cur
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
