package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题集：https://leetcode.cn/studyplan/top-interview-150/

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i, start, end := 0, -1, -1
	for i < len(intervals) {
		if intervals[i][1] >= newInterval[0] {
			start = i
			break
		} else {
			res = append(res, intervals[i])
		}
		i++
	}
	if start == -1 {
		res = append(res, newInterval)
		return res
	}

	for i < len(intervals) {
		if i == len(intervals)-1 || intervals[i+1][0] > newInterval[1] {
			end = i
			i++
			break
		}
		i++
	}
	if end == -1 {
		res = append(res, []int{intervals[start][0], newInterval[1]})
		return res
	}
	res = append(res, []int{intervals[start][0], max(intervals[end][1], newInterval[1])})
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func main() {
	// arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// res := averageOfLevels("abba", "dosg cat cat dog")
	// fmt.Println(res)
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{4, 5}))
}
