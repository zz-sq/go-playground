package main

import (
	"fmt"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i, j}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 3, 5, 6, 7}, 9))
}
