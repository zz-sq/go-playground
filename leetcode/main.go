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

func combine(n int, k int) [][]int {
	res := [][]int{}
	if k == 0 {
		return res
	}
	var backstrace func(j int, cur []int, visited []bool)
	backstrace = func(j int, cur []int, visited []bool) {
		if k == len(cur) {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := j; i <= n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			cur = append(cur, i)
			backstrace(i+1, cur, visited)
			visited[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	visited := make([]bool, n+1)
	backstrace(1, []int{}, visited)
	return res
}

func main() {
	// arr1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// res := averageOfLevels("abba", "dosg cat cat dog")
	// fmt.Println(res)
	fmt.Println(combine(1, 1))
}
