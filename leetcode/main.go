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

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	degree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		a, b := prerequisites[i][0], prerequisites[i][1]
		graph[b] = append(graph[b], a)
		degree[a]++
	}
	res := []int{}
	for i := 0; i < len(degree); i++ {
		if degree[i] == 0 {
			res = append(res, i)
		}
	}
	for i := 0; i < len(res); i++ {
		courses := graph[res[i]]
		for _, course := range courses {
			degree[course]--
			if degree[course] == 0 {
				res = append(res, course)
			}
		}
	}
	if len(res) == numCourses {
		return res
	}
	return []int{}
}

func main() {
	// arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// res := averageOfLevels("abba", "dosg cat cat dog")
	// fmt.Println(res)
	fmt.Println(findOrder(2, [][]int{{0, 1}}))
}
