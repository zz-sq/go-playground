package main

import (
	"fmt"
	"strings"
)

// 题集：https://leetcode.cn/studyplan/top-interview-150/
func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	res := make([]string, 0)
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != "" {
			res = append(res, strs[i])
		}
	}
	return strings.Join(res, " ")
}

// 1 2 3 4 5
// 3 4 5 1 2
// -2 -2 -2 4 3
func main() {
	fmt.Println(reverseWords("  hello world  "))
}
